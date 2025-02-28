package usecases

import (
	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/repositories"
)

type Usecase struct {
	AuthProfileUC AuthProfileUC
}

func NewUsecase(repo *repositories.Repository, cfg *config.Config) *Usecase {
	authProfileUC := NewAuthProfileUC(repo.AuthProfileRepo, repo.RefreshTokenRepo, cfg)
	return &Usecase{AuthProfileUC: authProfileUC}
}
