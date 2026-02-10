package user

import (
	"ms-practice/user-service/pkg/config"
	"ms-practice/user-service/pkg/usecase"
)

type userHandler struct {
	cfg    *config.Config
	userUC usecase.UserUC
}

func NewUserHandler(cfg *config.Config, uc usecase.Usecase) userHandler {
	return userHandler{
		cfg:    cfg,
		userUC: uc.UserUC,
	}
}
