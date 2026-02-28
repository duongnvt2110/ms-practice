package container

import (
	"fmt"
	"os"

	"ms-practice/pkg/db/gorm_client"
	"ms-practice/retry-management-service/pkg/config"
	"ms-practice/retry-management-service/pkg/repository"
	"ms-practice/retry-management-service/pkg/usecase"
)

type Container struct {
	Cfg   *config.Config
	DLQUC usecase.DLQUsecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()

	dbClient, err := gorm_client.NewGormClient(cfg.App.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	repo := repository.NewDLQRepository(dbClient)
	dlqUC := usecase.NewDLQUsecase(repo)

	return &Container{
		Cfg:   cfg,
		DLQUC: dlqUC,
	}
}
