package usecase

import (
	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/repository"
)

type Usecase struct {
	AuthProfileUC  AuthProfileUC
	UserGrpcClient UserGrpcClient
}

func NewUsecase(repo *repository.Repository, cfg *config.Config) *Usecase {
	userGrpcClient := NewUserGrpcClient(cfg)
	authProfileUC := NewAuthProfileUC(repo.AuthProfileRepo, repo.RefreshTokenRepo, userGrpcClient, cfg)
	return &Usecase{
		AuthProfileUC: authProfileUC,
	}
}
