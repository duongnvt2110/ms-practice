package repositories

import (
	"context"
	"ms-practice/auth-service/pkg/models"

	"gorm.io/gorm"
)

type RefreshTokenRepo interface {
	Create(ctx context.Context, rf *models.RefreshToken) error
	Delete(ctx context.Context, userID int, token string) error
}

type refreshTokenRepo struct {
	db *gorm.DB
}

var _ RefreshTokenRepo = (*refreshTokenRepo)(nil)

func NewRefreshTokenRepo(db *gorm.DB) RefreshTokenRepo {
	return &refreshTokenRepo{db: db}
}

func (r *refreshTokenRepo) Create(ctx context.Context, rf *models.RefreshToken) error {
	return r.db.WithContext(ctx).Create(rf).Error
}

func (r *refreshTokenRepo) Delete(ctx context.Context, userID int, token string) error {
	return r.db.WithContext(ctx).
		Where("user_id = ? and token = ?", userID, token).
		Delete(&models.RefreshToken{}).Error
}
