package container

import (
	"fmt"
	"ms-practice/booking-service/pkg/config"
	"ms-practice/booking-service/pkg/repository"
	"ms-practice/booking-service/pkg/usecase"
	"ms-practice/booking-service/pkg/util/kafka"
	"ms-practice/pkg/db/gorm_client"
	"os"
)

type Container struct {
	Cfg       *config.Config
	Usecases  *usecase.Usecase
	Messaging *kafka.BookingMessaging
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	bookingMessaging := kafka.NewBookingKafkaClient(cfg.App.Kafka)

	db, err := gorm_client.NewGormClient(cfg.App.Mysql)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	repo := repository.NewRepository(db)
	usecases := usecase.NewUsecase(repo, bookingMessaging)

	return &Container{
		Cfg:       cfg,
		Usecases:  usecases,
		Messaging: bookingMessaging,
	}
}

// Assumetion we have 1 milions requests -> create 1 milions connect to client if the send a message to a topic?
// - The issue occur
// Solution:
// Singleton for create client
