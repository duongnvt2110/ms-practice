package auth

import (
	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/usecases"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	cfg           *config.Config
	validate      *validator.Validate
	authProfileUC usecases.AuthProfileUC
}

func NewAuthHandler(cfg *config.Config, validate *validator.Validate, usecase *usecases.Usecase) AuthHandler {
	return AuthHandler{
		cfg:           cfg,
		authProfileUC: usecase.AuthProfileUC,
		validate:      validate,
	}
}
