package usecase

import (
	"ms-practice/user-service/pkg/config"
	"ms-practice/user-service/pkg/repository"
)

type Usecase struct {
	UserUC UserUC
	AuthUC AuthUC
}

func NewUsecase(
	repo *repository.Repository,
	cfg *config.Config) *Usecase {
	userUC := NewUserUC(repo.UserRepo)
	authUC := NewAuthUC(cfg)
	return &Usecase{
		UserUC: userUC,
		AuthUC: authUC,
	}
}
