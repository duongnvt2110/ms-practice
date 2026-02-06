package container

import (
	"fmt"
	"ms-practice/pkg/db/gorm_client"
	"ms-practice/user-service/pkg/config"
	"ms-practice/user-service/pkg/repository"
	"ms-practice/user-service/pkg/usecase"
	"os"
)

type Container struct {
	Cfg     *config.Config
	Usecase *usecase.Usecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	// Initialize dependencies
	dbClient, err := gorm_client.NewGormClient(cfg.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	repos := repository.NewRepository(dbClient)
	usecase := usecase.NewUsecase(repos, cfg)
	return &Container{
		Cfg:     cfg,
		Usecase: usecase,
	}
}
