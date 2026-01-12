package usecases

import (
	"context"
	"errors"
	"time"

	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/models"
	"ms-practice/auth-service/pkg/repositories"

	sharejwt "ms-practice/pkg/jwt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthProfileUC interface {
	Register(ctx context.Context, authProfileInfo *models.AuthProfile, userInfo *models.User) error
	Login(ctx context.Context, email, password string) (*models.TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (*models.TokenPair, error)
	Logout(ctx context.Context, token string) error
	ValidateToken(tokenString string) error
}

type authProfileUC struct {
	authRepo       repositories.AuthProfileRepo
	rfRepo         repositories.RefreshTokenRepo
	userGrpcClient UserGrpcClient
	cfg            *config.Config
}

var _ AuthProfileUC = (*authProfileUC)(nil)

func NewAuthProfileUC(authRepo repositories.AuthProfileRepo, rfRepo repositories.RefreshTokenRepo, userGrpcClient UserGrpcClient, cfg *config.Config) AuthProfileUC {
	return &authProfileUC{
		authRepo:       authRepo,
		rfRepo:         rfRepo,
		userGrpcClient: userGrpcClient,
		cfg:            cfg,
	}
}

// Public function
// Register a new user
func (u *authProfileUC) Register(ctx context.Context, authProfileInfo *models.AuthProfile, userInfo *models.User) error {
	// Check if user already exists
	existingUser, _ := u.authRepo.GetByEmail(ctx, authProfileInfo.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authProfileInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Grpc call create user in user-service
	userCreated, err := u.userGrpcClient.CreateUser(ctx, userInfo)
	if err != nil {
		return err
	}

	// Create auth_profile
	authProfileInfo.Password = string(hashedPassword)
	err = u.authRepo.Create(ctx, authProfileInfo)
	if err != nil {
		_, grpcError := u.userGrpcClient.DeleteUser(ctx, userCreated.GetId())
		if grpcError != nil {
			return grpcError
		}
		return err
	}
	// Todo
	// Implement verified
	// Send email
	return nil
}

// Login user and return JWT token
func (u *authProfileUC) Login(ctx context.Context, email, password string) (*models.TokenPair, error) {
	user, err := u.authRepo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := u.generateLoginToken(user.Email, user.Id)
	if err != nil {
		return nil, err
	}

	// enhance expires_time
	rf := &models.AuthRefreshToken{
		AuthProfileId: user.Id,
		RefreshToken:  token.RefreshToken,
		ExpiredAt:     time.Now().Add(time.Duration(u.cfg.JWT.RefreshTokenExp) * time.Minute),
	}
	err = u.rfRepo.Create(ctx, rf)
	if err != nil {
		return nil, err
	}

	return token, nil
}

func (u *authProfileUC) RefreshToken(ctx context.Context, refreshToken string) (*models.TokenPair, error) {
	if refreshToken == "" {
		return nil, errors.New("refresh token is required")
	}

	record, err := u.rfRepo.GetByToken(ctx, refreshToken)
	if err != nil {
		return nil, errors.New("invalid refresh token")
	}

	if time.Now().After(record.ExpiredAt) {
		_ = u.rfRepo.Delete(ctx, record.AuthProfileId, refreshToken)
		return nil, errors.New("refresh token expired")
	}

	claims := &models.AuthClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.cfg.JWT.Secret), nil
	})
	if err != nil || !token.Valid {
		return nil, errors.New("invalid refresh token")
	}

	newPair, err := u.generateLoginToken(claims.Email, claims.UserId)
	if err != nil {
		return nil, err
	}

	if err := u.rfRepo.Delete(ctx, record.AuthProfileId, refreshToken); err != nil {
		return nil, err
	}

	newRecord := &models.AuthRefreshToken{
		AuthProfileId: record.AuthProfileId,
		RefreshToken:  newPair.RefreshToken,
		ExpiredAt:     time.Now().Add(time.Duration(u.cfg.JWT.RefreshTokenExp) * time.Minute),
	}
	if err := u.rfRepo.Create(ctx, newRecord); err != nil {
		return nil, err
	}

	return newPair, nil
}

// Logout user (invalidate token)
func (u *authProfileUC) Logout(ctx context.Context, token string) error {
	// Here, you can implement token blacklisting if needed
	return nil
}

// ValidateJWT parses a token and checks its expiration
func (u *authProfileUC) ValidateToken(tokenString string) error {
	token, err := jwt.ParseWithClaims(tokenString, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(u.cfg.JWT.Secret), nil
	})

	if err != nil {
		return err
	}

	// Extract claims
	claims, ok := token.Claims.(*models.AuthClaims)
	if !ok || !token.Valid {
		return errors.New("invalid token")
	}

	// Check if the token is expired
	if claims.ExpiresAt.Before(time.Now()) {
		return errors.New("token expired")
	}

	return nil
}

// Private function
func (u *authProfileUC) generateLoginToken(email string, userID int) (*models.TokenPair, error) {
	accessExpiry := time.Now().Add(time.Duration(u.cfg.JWT.AccessTokenExp) * time.Minute)
	refreshExpiry := time.Now().Add(time.Duration(u.cfg.JWT.RefreshTokenExp) * time.Minute)

	// Create access token
	accessClaims := &models.AuthClaims{
		Email:  email,
		UserId: userID,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
		},
	}

	accessTokenString, err := sharejwt.JwtTokenEncode(u.cfg.JWT.Secret, accessClaims)
	if err != nil {
		return nil, err
	}

	// Create refresh token
	refreshClaims := models.AuthClaims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(refreshExpiry),
		},
	}

	refreshTokenString, err := sharejwt.JwtTokenEncode(u.cfg.JWT.Secret, refreshClaims)
	if err != nil {
		return nil, err
	}

	return &models.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
