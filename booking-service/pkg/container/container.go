package container

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"strconv"

	"ms-practice/booking-service/pkg/config"
	events "ms-practice/booking-service/pkg/event"
	"ms-practice/booking-service/pkg/model"
	"ms-practice/booking-service/pkg/repository"
	usecase "ms-practice/booking-service/pkg/usecase"
	bookingUsecase "ms-practice/booking-service/pkg/usecase/booking"
	"ms-practice/booking-service/pkg/util/kafka"
	"ms-practice/pkg/db/gorm_client"
	sharedKafka "ms-practice/pkg/kafka"
)

type Container struct {
	Cfg              *config.Config
	BookingMessaging *kafka.BookingMessaging
	Usecases         *usecase.Usecase
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()

	kafkaClient := sharedKafka.NewKafkaClient(cfg.App.Kafka)
	bookingMessaging := kafka.NewBookingKafkaClient(kafkaClient)

	dbClient, err := gorm_client.NewGormClient(cfg.App.Mysql.PrimaryHosts, cfg.App.Mysql.ReadHosts, cfg.App.Mysql.User, cfg.App.Mysql.Password, cfg.App.Mysql.Port, cfg.App.Mysql.DBName)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	repos := repository.NewRepository(dbClient)
	publisher := newKafkaBookingPublisher(bookingMessaging)
	usecases := usecase.NewUsecase(repos, publisher)

	return &Container{
		Cfg:              cfg,
		BookingMessaging: bookingMessaging,
		Usecases:         usecases,
	}
}

type kafkaBookingPublisher struct {
	messaging *kafka.BookingMessaging
}

func newKafkaBookingPublisher(messaging *kafka.BookingMessaging) bookingUsecase.BookingPublisher {
	return &kafkaBookingPublisher{messaging: messaging}
}

func (p *kafkaBookingPublisher) BookingCreated(ctx context.Context, booking *model.Booking) error {
	if p.messaging == nil {
		return nil
	}
	producer, ok := p.messaging.Producers[events.BookingTopicName]
	if !ok || producer == nil {
		return nil
	}

	payload := events.BookingOrdered{
		EventType: "BookingOrdered",
		OrderID:   strconv.Itoa(booking.Id),
		Amount:    booking.Prices,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return producer.Publish(ctx, nil, data)
}

// Assumetion we have 1 milions requests -> create 1 milions connect to client if the send a message to a topic?
// - The issue occur
// Solution:
// Singleton for create client
