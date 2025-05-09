package user

import (
	"ms-practice/user-service/pkg/config"
	"ms-practice/user-service/pkg/usecases"
)

type userHandler struct {
	cfg    *config.Config
	userUC usecases.UserUsecase
}

func NewUserHandler(cfg *config.Config, uc usecases.Usecase) userHandler {
	return userHandler{
		cfg:    cfg,
		userUC: uc.UserUC,
	}
}
