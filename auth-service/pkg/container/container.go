package container

import (
	"auth-service/pkg/config"
)

type Container struct {
	Cfg *config.Config
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	return &Container{
		Cfg: cfg,
	}
}
