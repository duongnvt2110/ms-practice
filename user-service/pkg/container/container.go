package container

import (
	"fmt"
	"ms-practice/pkg/db/gorm_client"
	"ms-practice/user-service/pkg/config"
	"ms-practice/user-service/pkg/repositories"
	"ms-practice/user-service/pkg/usecases"
	"os"
)

type Container struct {
	Cfg     *config.Config
	Usecase *usecases.Usecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	// Initialize dependencies
	dbClient, err := gorm_client.NewGormClient(cfg.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	repos := repositories.NewRepository(dbClient)
	usecases := usecases.NewUsecase(repos, cfg)
	return &Container{
		Cfg:     cfg,
		Usecase: usecases,
	}
}
