package container

import (
	"fmt"
	"ms-practice/payment-service/pkg/config"
	"ms-practice/payment-service/pkg/repository"
	"ms-practice/payment-service/pkg/usecase"
	"ms-practice/pkg/db/gorm_client"
	"os"
)

type Container struct {
	Cfg      *config.Config
	Usecases *usecase.Usecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	// Initialize dependencies
	db, err := gorm_client.NewGormClient(cfg.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	repo := repository.NewRepository(db)
	usecases := usecase.NewUsecase(repo)
	return &Container{
		Cfg:      cfg,
		Usecases: usecases,
	}
}
