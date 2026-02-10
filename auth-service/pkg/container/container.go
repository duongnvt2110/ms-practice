package container

import (
	"fmt"
	"os"

	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/repository"
	"ms-practice/auth-service/pkg/usecase"

	"ms-practice/pkg/db/gorm_client"
	svalidate "ms-practice/pkg/validator"

	"github.com/go-playground/validator/v10"
)

type Container struct {
	Cfg      *config.Config
	Usecase  *usecase.Usecase
	Validate *validator.Validate
}

func InitializeContainer() *Container {
	// Initialize dependencies
	cfg := config.NewConfig()
	validate := svalidate.NewValidate()
	dbClient, err := gorm_client.NewGormClient(cfg.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	repo := repository.NewRepository(dbClient)
	usecase := usecase.NewUsecase(repo, cfg)

	return &Container{
		Cfg:      cfg,
		Usecase:  usecase,
		Validate: validate,
	}
}
