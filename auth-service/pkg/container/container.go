package container

import (
	"auth-service/pkg/config"
	"auth-service/pkg/repositories"
	"auth-service/pkg/usecases"
	"fmt"
	"os"

	"pkg/db/gorm_client"
)

type Container struct {
	Cfg     *config.Config
	Usecase *usecases.Usecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	dbClient, err := gorm_client.NewGormClient(cfg.Mysql.PrimaryHosts, cfg.Mysql.ReadHosts, cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Port, cfg.Mysql.DBName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	repo := repositories.NewRepository(dbClient)
	usecase := usecases.NewUsecase(repo, cfg)
	return &Container{
		Cfg:     cfg,
		Usecase: usecase,
	}
}
