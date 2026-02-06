package auth

import (
	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/usecase"

	"github.com/go-playground/validator/v10"
)

type AuthHandler struct {
	cfg           *config.Config
	validate      *validator.Validate
	authProfileUC usecase.AuthProfileUC
}

func NewAuthHandler(cfg *config.Config, validate *validator.Validate, usecase *usecase.Usecase) AuthHandler {
	return AuthHandler{
		cfg:           cfg,
		authProfileUC: usecase.AuthProfileUC,
		validate:      validate,
	}
}
