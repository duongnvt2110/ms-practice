package container

import (
	"fmt"
	"os"

	"ms-practice/auth-service/pkg/config"
	"ms-practice/auth-service/pkg/repositories"
	"ms-practice/auth-service/pkg/usecases"

	"ms-practice/pkg/db/gorm_client"
	svalidate "ms-practice/pkg/validator"

	"github.com/go-playground/validator/v10"
)

type Container struct {
	Cfg      *config.Config
	Usecase  *usecases.Usecase
	Validate *validator.Validate
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	validate := svalidate.NewValidate()
	dbClient, err := gorm_client.NewGormClient(cfg.Mysql.PrimaryHosts, cfg.Mysql.ReadHosts, cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Port, cfg.Mysql.DBName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	repo := repositories.NewRepository(dbClient)
	usecase := usecases.NewUsecase(repo, cfg)
	return &Container{
		Cfg:      cfg,
		Usecase:  usecase,
		Validate: validate,
	}
}
