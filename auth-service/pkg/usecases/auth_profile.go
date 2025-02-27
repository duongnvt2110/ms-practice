package usecases

import (
	"auth-service/pkg/config"
	"auth-service/pkg/models"
	"auth-service/pkg/repositories"
	"context"
	"errors"
	"time"

	sharejwt "pkg/jwt"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type AuthProfileUC interface {
	Register(ctx context.Context, authProfileInfo *models.AuthProfile, userInfo *models.User) error
	Login(ctx context.Context, email, password string) (*models.TokenPair, error)
	Logout(ctx context.Context, token string) error
	ValidateToken(tokenString string) error
}

type authProfileUC struct {
	repo repositories.AuthProfileRepo
	cfg  *config.Config
}

var _ AuthProfileUC = (*authProfileUC)(nil)

func NewAuthProfileUC(repo repositories.AuthProfileRepo, cfg *config.Config) AuthProfileUC {
	return &authProfileUC{repo: repo, cfg: cfg}
}

// Public function
// Register a new user
func (u *authProfileUC) Register(ctx context.Context, authProfileInfo *models.AuthProfile, userInfo *models.User) error {
	// Check if user already exists
	existingUser, _ := u.repo.GetByEmail(ctx, authProfileInfo.Email)
	if existingUser != nil {
		return errors.New("user already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(authProfileInfo.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// Create new user
	authProfileInfo.Password = string(hashedPassword)

	return u.repo.Create(ctx, authProfileInfo)

	// Todo
	// Grpc call create user in user-service
	// Send email
}

// Login user and return JWT token
func (u *authProfileUC) Login(ctx context.Context, email, password string) (*models.TokenPair, error) {
	user, err := u.repo.GetByEmail(ctx, email)
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Compare hashed password
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := u.generateLoginToken(email)
	if err != nil {
		return nil, err
	}

	return token, nil
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
func (u *authProfileUC) generateLoginToken(email string) (*models.TokenPair, error) {
	accessExpiry := time.Now().Add(time.Duration(u.cfg.JWT.AccessTokenExp) * time.Minute)
	refreshExpiry := time.Now().Add(time.Duration(u.cfg.JWT.RefreshTokenExp) * time.Minute)

	// Create access token
	accessClaims := &models.AuthClaims{
		Email: email,
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
