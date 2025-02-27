package usecases

import (
	"auth-service/pkg/config"
	"auth-service/pkg/repositories"
)

type Usecase struct {
	AuthProfileUC AuthProfileUC
}

func NewUsecase(repo *repositories.Repository, cfg *config.Config) *Usecase {
	authProfileUC := NewAuthProfileUC(repo.AuthProfileRepo, cfg)
	return &Usecase{AuthProfileUC: authProfileUC}
}
