package usecases

import (
	"context"
	"ms-practice/user-service/pkg/models"
	"ms-practice/user-service/pkg/repositories"
)

type UserUsecase interface {
	GetUser(ctx context.Context, id int32) (*models.User, error)
}

type userUsecase struct {
	userRepo repositories.UserRepository
}

func NewUserUsecase(repo repositories.UserRepository) UserUsecase {
	return &userUsecase{userRepo: repo}
}

func (u *userUsecase) GetUser(ctx context.Context, id int32) (*models.User, error) {
	return u.userRepo.GetUserByID(ctx, id)
}
