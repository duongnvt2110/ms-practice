package usecase

import (
	"context"
	"ms-practice/user-service/pkg/model"
	"ms-practice/user-service/pkg/repository"
)

type UserUC interface {
	GetUser(ctx context.Context, id int32) (*model.User, error)
	CreateUser(ctx context.Context, user *model.User) (int32, error)
	DeleteUser(ctx context.Context, userID int32) error
}

type userUC struct {
	userRepo repository.UserRepository
}

func NewUserUC(repo repository.UserRepository) UserUC {
	return &userUC{userRepo: repo}
}

func (uc *userUC) GetUser(ctx context.Context, id int32) (*model.User, error) {
	return uc.userRepo.GetUserByID(ctx, id)
}

func (uc *userUC) CreateUser(ctx context.Context, user *model.User) (int32, error) {
	return uc.userRepo.CreateUser(ctx, user)
}

func (uc *userUC) DeleteUser(ctx context.Context, userID int32) error {
	return uc.userRepo.DeleteUser(ctx, userID)
}
