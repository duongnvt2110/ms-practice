package repositories

import (
	"context"
	"errors"

	"ms-practice/auth-service/pkg/models"

	"gorm.io/gorm"
)

type AuthProfileRepo interface {
	Create(ctx context.Context, user *models.AuthProfile) error
	GetByEmail(ctx context.Context, email string) (*models.AuthProfile, error)
	Update(ctx context.Context, user *models.AuthProfile) error
	Delete(ctx context.Context, email string) error
}

type authProfileRepo struct {
	db *gorm.DB
}

var _ AuthProfileRepo = (*authProfileRepo)(nil)

func NewAuthProfileRepo(db *gorm.DB) AuthProfileRepo {
	return &authProfileRepo{db: db}
}

// Create a new user
func (r *authProfileRepo) Create(ctx context.Context, user *models.AuthProfile) error {
	return r.db.WithContext(ctx).Create(user).Error
}

// Get user by email
func (r *authProfileRepo) GetByEmail(ctx context.Context, email string) (*models.AuthProfile, error) {
	var user models.AuthProfile
	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

// Update user details
func (r *authProfileRepo) Update(ctx context.Context, user *models.AuthProfile) error {
	return r.db.WithContext(ctx).Save(user).Error
}

// Delete user by email
func (r *authProfileRepo) Delete(ctx context.Context, email string) error {
	return r.db.WithContext(ctx).Where("email = ?", email).Delete(&models.AuthProfile{}).Error
}
