package usecases

import (
	"context"
	"errors"
	"testing"
	"time"

	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/models"
	autherror "ms-practice/auth-service/pkg/utils/errors"
	apperror "ms-practice/pkg/errors"
	sharejwt "ms-practice/pkg/jwt"
	"ms-practice/proto/gen"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"google.golang.org/protobuf/types/known/emptypb"
)

type fakeAuthProfileRepo struct {
	getByEmailFn func(ctx context.Context, email string) (*models.AuthProfile, error)
	createFn     func(ctx context.Context, user *models.AuthProfile) error
	updateFn     func(ctx context.Context, user *models.AuthProfile) error
	deleteFn     func(ctx context.Context, email string) error
}

func (f *fakeAuthProfileRepo) Create(ctx context.Context, user *models.AuthProfile) error {
	if f.createFn != nil {
		return f.createFn(ctx, user)
	}
	return nil
}

func (f *fakeAuthProfileRepo) GetByEmail(ctx context.Context, email string) (*models.AuthProfile, error) {
	if f.getByEmailFn != nil {
		return f.getByEmailFn(ctx, email)
	}
	return nil, nil
}

func (f *fakeAuthProfileRepo) Update(ctx context.Context, user *models.AuthProfile) error {
	if f.updateFn != nil {
		return f.updateFn(ctx, user)
	}
	return nil
}

func (f *fakeAuthProfileRepo) Delete(ctx context.Context, email string) error {
	if f.deleteFn != nil {
		return f.deleteFn(ctx, email)
	}
	return nil
}

type fakeRefreshTokenRepo struct {
	createFn   func(ctx context.Context, rf *models.AuthRefreshToken) error
	deleteFn   func(ctx context.Context, authProfileID int, token string) error
	getByToken func(ctx context.Context, token string) (*models.AuthRefreshToken, error)
}

func (f *fakeRefreshTokenRepo) Create(ctx context.Context, rf *models.AuthRefreshToken) error {
	if f.createFn != nil {
		return f.createFn(ctx, rf)
	}
	return nil
}

func (f *fakeRefreshTokenRepo) Delete(ctx context.Context, authProfileID int, token string) error {
	if f.deleteFn != nil {
		return f.deleteFn(ctx, authProfileID, token)
	}
	return nil
}

func (f *fakeRefreshTokenRepo) GetByToken(ctx context.Context, token string) (*models.AuthRefreshToken, error) {
	if f.getByToken != nil {
		return f.getByToken(ctx, token)
	}
	return nil, nil
}

type fakeUserGrpcClient struct {
	createUserFn func(ctx context.Context, user *models.User) (*gen.CreateUserResponse, error)
	deleteUserFn func(ctx context.Context, userID int32) (*emptypb.Empty, error)
}

func (f *fakeUserGrpcClient) GetUser(ctx context.Context, id int32) (*gen.User, error) {
	return nil, nil
}

func (f *fakeUserGrpcClient) CreateUser(ctx context.Context, user *models.User) (*gen.CreateUserResponse, error) {
	if f.createUserFn != nil {
		return f.createUserFn(ctx, user)
	}
	return &gen.CreateUserResponse{}, nil
}

func (f *fakeUserGrpcClient) DeleteUser(ctx context.Context, userID int32) (*emptypb.Empty, error) {
	if f.deleteUserFn != nil {
		return f.deleteUserFn(ctx, userID)
	}
	return &emptypb.Empty{}, nil
}

func newTestConfig() *config.Config {
	cfg := &config.Config{}
	cfg.JWT.Secret = "test-secret"
	cfg.JWT.AccessTokenExp = 10
	cfg.JWT.RefreshTokenExp = 10
	return cfg
}

func assertAppError(t *testing.T, err error, expected apperror.AppError) {
	t.Helper()
	if err == nil {
		t.Fatalf("expected error, got nil")
	}
	var appErr apperror.AppError
	if !errors.As(err, &appErr) {
		t.Fatalf("expected AppError, got %T", err)
	}
	if appErr.GetErrCode() != expected.GetErrCode() {
		t.Fatalf("expected err code %s, got %s", expected.GetErrCode(), appErr.GetErrCode())
	}
	if appErr.PublicMessage() != expected.PublicMessage() {
		t.Fatalf("expected message %q, got %q", expected.PublicMessage(), appErr.PublicMessage())
	}
}

func TestRegister_ErrorMapping(t *testing.T) {
	cfg := newTestConfig()
	tests := []struct {
		name    string
		repo    *fakeAuthProfileRepo
		wantErr apperror.AppError
	}{
		{
			name: "user already exists",
			repo: &fakeAuthProfileRepo{
				getByEmailFn: func(ctx context.Context, email string) (*models.AuthProfile, error) {
					return &models.AuthProfile{Email: email}, nil
				},
			},
			wantErr: autherror.ErrUserAlreadyExists,
		},
		{
			name: "repo error",
			repo: &fakeAuthProfileRepo{
				getByEmailFn: func(ctx context.Context, email string) (*models.AuthProfile, error) {
					return nil, errors.New("db down")
				},
			},
			wantErr: apperror.ErrInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewAuthProfileUC(tt.repo, &fakeRefreshTokenRepo{}, &fakeUserGrpcClient{}, cfg)
			err := uc.Register(context.Background(), &models.AuthProfile{Email: "a@b.com", Password: "pw"}, &models.User{})
			assertAppError(t, err, tt.wantErr)
		})
	}
}

func TestLogin_ErrorMapping(t *testing.T) {
	cfg := newTestConfig()
	hashed, _ := bcrypt.GenerateFromPassword([]byte("right"), bcrypt.DefaultCost)
	tests := []struct {
		name    string
		repo    *fakeAuthProfileRepo
		email   string
		pass    string
		wantErr apperror.AppError
	}{
		{
			name: "user not found",
			repo: &fakeAuthProfileRepo{
				getByEmailFn: func(ctx context.Context, email string) (*models.AuthProfile, error) {
					return nil, autherror.ErrUserNotFound
				},
			},
			email:   "a@b.com",
			pass:    "pw",
			wantErr: autherror.ErrInvalidCredentials,
		},
		{
			name: "invalid password",
			repo: &fakeAuthProfileRepo{
				getByEmailFn: func(ctx context.Context, email string) (*models.AuthProfile, error) {
					return &models.AuthProfile{Email: email, Password: string(hashed)}, nil
				},
			},
			email:   "a@b.com",
			pass:    "wrong",
			wantErr: autherror.ErrInvalidCredentials,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewAuthProfileUC(tt.repo, &fakeRefreshTokenRepo{}, &fakeUserGrpcClient{}, cfg)
			_, err := uc.Login(context.Background(), tt.email, tt.pass)
			assertAppError(t, err, tt.wantErr)
		})
	}
}

func TestRefreshToken_ErrorMapping(t *testing.T) {
	cfg := newTestConfig()
	tests := []struct {
		name       string
		repo       *fakeRefreshTokenRepo
		token      string
		wantErr    apperror.AppError
		setupToken bool
	}{
		{
			name:    "refresh token required",
			token:   "",
			wantErr: autherror.ErrRefreshTokenRequired,
		},
		{
			name: "invalid refresh token",
			repo: &fakeRefreshTokenRepo{
				getByToken: func(ctx context.Context, token string) (*models.AuthRefreshToken, error) {
					return nil, autherror.ErrInvalidRefreshToken
				},
			},
			token:   "bad",
			wantErr: autherror.ErrInvalidRefreshToken,
		},
		{
			name: "refresh token expired",
			repo: &fakeRefreshTokenRepo{
				getByToken: func(ctx context.Context, token string) (*models.AuthRefreshToken, error) {
					return &models.AuthRefreshToken{
						AuthProfileId: 1,
						RefreshToken:  token,
						ExpiredAt:     time.Now().Add(-time.Hour),
					}, nil
				},
			},
			token:   "expired",
			wantErr: autherror.ErrRefreshTokenExpired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := tt.repo
			if repo == nil {
				repo = &fakeRefreshTokenRepo{}
			}
			uc := NewAuthProfileUC(&fakeAuthProfileRepo{}, repo, &fakeUserGrpcClient{}, cfg)
			_, err := uc.RefreshToken(context.Background(), tt.token)
			assertAppError(t, err, tt.wantErr)
		})
	}
}

func TestValidateToken_ErrorMapping(t *testing.T) {
	cfg := newTestConfig()
	expiredClaims := &models.AuthClaims{
		Email: "a@b.com",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(-time.Minute)),
		},
	}
	expiredToken, err := sharejwt.JwtTokenEncode(cfg.JWT.Secret, expiredClaims)
	if err != nil {
		t.Fatalf("failed to create token: %v", err)
	}

	tests := []struct {
		name    string
		token   string
		wantErr apperror.AppError
	}{
		{
			name:    "invalid token",
			token:   "bad-token",
			wantErr: autherror.ErrInvalidToken,
		},
		{
			name:    "token expired",
			token:   expiredToken,
			wantErr: autherror.ErrTokenExpired,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			uc := NewAuthProfileUC(&fakeAuthProfileRepo{}, &fakeRefreshTokenRepo{}, &fakeUserGrpcClient{}, cfg)
			_, err := uc.ValidateToken(tt.token)
			assertAppError(t, err, tt.wantErr)
		})
	}
}
