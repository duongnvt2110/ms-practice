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
	Cfg    *config.Config
	UserUc usecases.UserUsecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	// Initialize dependencies
	dbClient, err := gorm_client.NewGormClient(cfg.Mysql.PrimaryHosts, cfg.Mysql.ReadHosts, cfg.Mysql.User, cfg.Mysql.Password, cfg.Mysql.Port, cfg.Mysql.DBName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	userRepo := repositories.NewUserRepository(dbClient)
	userUC := usecases.NewUserUsecase(userRepo)
	return &Container{
		Cfg:    cfg,
		UserUc: userUC,
	}
}
