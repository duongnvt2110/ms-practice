package usecase

import (
	"context"
	"errors"
	"time"

	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/models"
	"ms-practice/auth-service/pkg/repository"
	"ms-practice/auth-service/pkg/utils/apperr"

	"ms-practice/pkg/errorsx"
	sharejwt "ms-practice/pkg/jwt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthProfileUC interface {
	Register(ctx context.Context, authProfileInfo *models.AuthProfile, userInfo *models.User) error
	Login(ctx context.Context, email, password string) (*models.TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (*models.TokenPair, error)
	Logout(ctx context.Context, authProfileID int, token string) error
	ValidateToken(tokenString string) (*models.AuthClaims, error)
}

type authProfileUC struct {
	authRepo       repository.AuthProfileRepo
	rfRepo         repository.RefreshTokenRepo
	userGrpcClient UserGrpcClient
	cfg            *config.Config
}

var _ AuthProfileUC = (*authProfileUC)(nil)

func NewAuthProfileUC(authRepo repository.AuthProfileRepo, rfRepo repository.RefreshTokenRepo, userGrpcClient UserGrpcClient, cfg *config.Config) AuthProfileUC {
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
	existingUser, err := u.authRepo.GetByEmail(ctx, authProfileInfo.Email)
	if err != nil {
		var appErr errorsx.AppError
		if errors.As(err, &appErr) && appErr.GetErrCode() == apperr.ErrUserNotFound.GetErrCode() {
			err = nil
		} else {
			return errorsx.ErrInternalServerError.Wrap(err)
		}
	}
	if existingUser != nil {
		return apperr.ErrUserAlreadyExists
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authProfileInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return errorsx.ErrInternalServerError.Wrap(err)
	}

	// Grpc call create user in user-service
	userCreated, err := u.userGrpcClient.CreateUser(ctx, userInfo)
	if err != nil {
		return errorsx.ErrInternalServerError.Wrap(err)
	}

	// Create auth_profile
	authProfileInfo.Password = string(hashedPassword)
	err = u.authRepo.Create(ctx, authProfileInfo)
	if err != nil {
		_, grpcError := u.userGrpcClient.DeleteUser(ctx, userCreated.GetId())
		if grpcError != nil {
			return errorsx.ErrInternalServerError.Wrap(grpcError)
		}
		return errorsx.ErrInternalServerError.Wrap(err)
	}
	// Todo
	// Implement verified
	// Send email
	return nil
}

// Login user and return JWT token
func (u *authProfileUC) Login(ctx context.Context, email, password string) (*models.TokenPair, error) {
	authProfile, err := u.authRepo.GetByEmail(ctx, email)
	if err != nil {
		var appErr errorsx.AppError
		if errors.As(err, &appErr) {
			if appErr.GetErrCode() == apperr.ErrUserNotFound.GetErrCode() {
				return nil, apperr.ErrInvalidCredentials
			}
			return nil, appErr
		}
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(authProfile.Password), []byte(password))
	if err != nil {
		return nil, apperr.ErrInvalidCredentials
	}

	// Generate JWT token
	token, err := u.generateLoginToken(authProfile.Email, authProfile.Id)
	if err != nil {
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	// enhance expires_time
	rf := &models.AuthRefreshToken{
		AuthProfileId: authProfile.Id,
		RefreshToken:  token.RefreshToken,
		ExpiredAt:     time.Now().Add(time.Duration(u.cfg.JWT.RefreshTokenExp) * time.Minute),
	}
	err = u.rfRepo.Create(ctx, rf)
	if err != nil {
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	return token, nil
}

func (u *authProfileUC) RefreshToken(ctx context.Context, refreshToken string) (*models.TokenPair, error) {
	if refreshToken == "" {
		return nil, apperr.ErrRefreshTokenRequired
	}

	record, err := u.rfRepo.GetByToken(ctx, refreshToken)
	if err != nil {
		var appErr errorsx.AppError
		if errors.As(err, &appErr) {
			return nil, appErr
		}
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	if time.Now().After(record.ExpiredAt) {
		_ = u.rfRepo.Delete(ctx, record.AuthProfileId, refreshToken)
		return nil, apperr.ErrRefreshTokenExpired
	}

	claims := &models.AuthClaims{}
	token, err := jwt.ParseWithClaims(refreshToken, claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperr.ErrInvalidRefreshToken
		}
		return []byte(u.cfg.JWT.Secret), nil
	})
	if err != nil || !token.Valid {
		return nil, apperr.ErrInvalidRefreshToken
	}

	newPair, err := u.generateLoginToken(claims.Email, claims.AuthProfileId)
	if err != nil {
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	if err := u.rfRepo.Delete(ctx, record.AuthProfileId, refreshToken); err != nil {
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	newRecord := &models.AuthRefreshToken{
		AuthProfileId: record.AuthProfileId,
		RefreshToken:  newPair.RefreshToken,
		ExpiredAt:     time.Now().Add(time.Duration(u.cfg.JWT.RefreshTokenExp) * time.Minute),
	}
	if err := u.rfRepo.Create(ctx, newRecord); err != nil {
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	return newPair, nil
}

// Logout user (invalidate token)
func (u *authProfileUC) Logout(ctx context.Context, authProfileID int, token string) error {
	// Here, you can implement token blacklisting if needed
	return u.rfRepo.Delete(ctx, authProfileID, token)
}

// ValidateJWT parses a token and checks its expiration
func (u *authProfileUC) ValidateToken(tokenString string) (*models.AuthClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &models.AuthClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, apperr.ErrInvalidToken
		}
		return []byte(u.cfg.JWT.Secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, apperr.ErrTokenExpired
		}
		return nil, apperr.ErrInvalidToken
	}

	// Extract claims
	claims, ok := token.Claims.(*models.AuthClaims)
	if !ok || !token.Valid {
		return nil, apperr.ErrInvalidToken
	}

	// Check if the token is expired
	if claims.ExpiresAt.Before(time.Now()) {
		return nil, apperr.ErrTokenExpired
	}

	return claims, nil
}

// Private function
func (u *authProfileUC) generateLoginToken(email string, authProfileId int) (*models.TokenPair, error) {
	accessExpiry := time.Now().Add(time.Duration(u.cfg.JWT.AccessTokenExp) * time.Minute)
	refreshExpiry := time.Now().Add(time.Duration(u.cfg.JWT.RefreshTokenExp) * time.Minute)

	// Create access token
	accessClaims := &models.AuthClaims{
		Email:         email,
		AuthProfileId: authProfileId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(accessExpiry),
		},
	}

	accessTokenString, err := sharejwt.JwtTokenEncode(u.cfg.JWT.Secret, accessClaims)
	if err != nil {
		return nil, errorsx.ErrInternalServerError.Wrap(err)
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
		return nil, errorsx.ErrInternalServerError.Wrap(err)
	}

	return &models.TokenPair{
		AccessToken:  accessTokenString,
		RefreshToken: refreshTokenString,
	}, nil
}
