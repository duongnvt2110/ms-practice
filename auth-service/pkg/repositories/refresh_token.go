package repositories

import (
	"context"
	"ms-practice/auth-service/pkg/models"

	"gorm.io/gorm"
)

type RefreshTokenRepo interface {
	Create(ctx context.Context, rf *models.AuthRefreshToken) error
	Delete(ctx context.Context, authProfileID int, token string) error
	GetByToken(ctx context.Context, token string) (*models.AuthRefreshToken, error)
}

type refreshTokenRepo struct {
	db *gorm.DB
}

var _ RefreshTokenRepo = (*refreshTokenRepo)(nil)

func NewRefreshTokenRepo(db *gorm.DB) RefreshTokenRepo {
	return &refreshTokenRepo{db: db}
}

func (r *refreshTokenRepo) Create(ctx context.Context, rf *models.AuthRefreshToken) error {
	return r.db.WithContext(ctx).Create(rf).Error
}

func (r *refreshTokenRepo) Delete(ctx context.Context, authProfileID int, token string) error {
	return r.db.WithContext(ctx).
		Where("auth_profile_id = ? AND refresh_token = ?", authProfileID, token).
		Delete(&models.AuthRefreshToken{}).Error
}

func (r *refreshTokenRepo) GetByToken(ctx context.Context, token string) (*models.AuthRefreshToken, error) {
	var refreshToken models.AuthRefreshToken
	if err := r.db.WithContext(ctx).
		Where("refresh_token = ?", token).
		First(&refreshToken).Error; err != nil {
		return nil, err
	}
	return &refreshToken, nil
}
