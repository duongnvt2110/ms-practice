package usecases

import (
	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/repositories"
)

type Usecase struct {
	AuthProfileUC  AuthProfileUC
	UserGrpcClient UserGrpcClient
}

func NewUsecase(repo *repositories.Repository, cfg *config.Config) *Usecase {
	userGrpcClient := NewUserGrpcClient(cfg)
	authProfileUC := NewAuthProfileUC(repo.AuthProfileRepo, repo.RefreshTokenRepo, userGrpcClient, cfg)
	return &Usecase{
		AuthProfileUC: authProfileUC,
	}
}
