package repositories

import (
	"context"
	"ms-practice/user-service/pkg/models"

	"gorm.io/gorm"
)

type UserRepository interface {
	GetUserByID(ctx context.Context, id int32) (*models.User, error)
	CreateUser(ctx context.Context, user *models.User) (int32, error)
	DeleteUser(ctx context.Context, userID int32) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) GetUserByID(ctx context.Context, id int32) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).
		Preload("Settings").
		Where("id = ?", id).
		First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) CreateUser(ctx context.Context, user *models.User) (int32, error) {
	err := r.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return 0, err
	}
	return user.Id, nil
}

func (r *userRepository) DeleteUser(ctx context.Context, userID int32) error {
	return r.db.WithContext(ctx).Delete(&models.User{}, userID).Error
}
