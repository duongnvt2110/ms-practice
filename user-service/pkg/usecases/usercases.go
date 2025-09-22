package usecases

import (
	"ms-practice/user-service/pkg/config"
	"ms-practice/user-service/pkg/repositories"
)

type Usecase struct {
	UserUC UserUsecase
}

func NewUsecase(repo *repositories.Repository, cfg *config.Config) *Usecase {
	userUsecase := NewUserUsecase(repo.UserRepo)
	return &Usecase{
		UserUC: userUsecase,
	}
}
