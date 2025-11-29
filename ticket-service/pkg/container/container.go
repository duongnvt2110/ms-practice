package container

import (
	"fmt"
	"ms-practice/pkg/db/gorm_client"
	"ms-practice/ticket-service/pkg/config"
	"ms-practice/ticket-service/pkg/repositories"
	"ms-practice/ticket-service/pkg/usecase"
	"os"
)

type Container struct {
	Cfg      *config.Config
	Usecases *usecase.Usecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()

	db, err := gorm_client.NewGormClient(cfg.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	repos := repositories.NewRepository(db)
	usecases := usecase.NewUsecase(repos)

	return &Container{
		Cfg:      cfg,
		Usecases: usecases,
	}
}
