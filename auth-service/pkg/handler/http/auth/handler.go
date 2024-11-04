package auth

import "auth-service/pkg/config"

type authHandler struct {
	cfg *config.Config
}

func NewAuthHandler(cfg *config.Config) authHandler {
	return authHandler{
		cfg: cfg,
	}
}
