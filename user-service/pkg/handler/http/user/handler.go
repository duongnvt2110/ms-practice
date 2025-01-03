package user

import "user-service/pkg/config"

type userHandler struct {
	cfg *config.Config
}

func NewUserHandler(cfg *config.Config) userHandler {
	return userHandler{
		cfg: cfg,
	}
}
