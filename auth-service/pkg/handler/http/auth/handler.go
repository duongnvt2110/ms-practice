package auth

import (
	"auth-service/pkg/config"
	"auth-service/pkg/usecases"
)

type AuthHandler struct {
	cfg           *config.Config
	authProfileUC usecases.AuthProfileUC
}

func NewAuthHandler(cfg *config.Config, usecase *usecases.Usecase) AuthHandler {
	return AuthHandler{
		cfg:           cfg,
		authProfileUC: usecase.AuthProfileUC,
	}
}
