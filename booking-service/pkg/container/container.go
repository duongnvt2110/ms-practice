package container

import (
	"fmt"
	"ms-practice/booking-service/pkg/config"
	"ms-practice/booking-service/pkg/repository"
	"ms-practice/booking-service/pkg/usecase"
	"ms-practice/booking-service/pkg/util/kafka"
	"ms-practice/pkg/db/gorm_client"
	sharedKafka "ms-practice/pkg/kafka"
	"os"
)

type Container struct {
	Cfg      *config.Config
	Usecases *usecase.Usecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	kafkaClient := sharedKafka.NewKafkaClient(cfg.App.Kafka)
	bookingMessaging := kafka.NewBookingKafkaClient(kafkaClient)

	db, err := gorm_client.NewGormClient(cfg.App.Mysql.PrimaryHosts, cfg.App.Mysql.ReadHosts, cfg.App.Mysql.User, cfg.App.Mysql.Password, cfg.App.Mysql.Port, cfg.App.Mysql.DBName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	repo := repository.NewRepository(db)
	usecases := usecase.NewUsecase(repo, bookingMessaging)

	return &Container{
		Cfg:      cfg,
		Usecases: usecases,
	}
}

// Assumetion we have 1 milions requests -> create 1 milions connect to client if the send a message to a topic?
// - The issue occur
// Solution:
// Singleton for create client
