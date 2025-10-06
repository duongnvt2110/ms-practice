package container

import (
	"ms-practice/booking-service/pkg/config"
	"ms-practice/booking-service/pkg/util/kafka"
	sharedKafka "ms-practice/pkg/kafka"
)

type Container struct {
	Cfg              *config.Config
	BookingMessaging *kafka.BookingMessaging
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	kafkaClient := sharedKafka.NewKafkaClient(cfg.App.Kafka)
	k := kafka.NewBookingKafkaClient(kafkaClient)
	return &Container{
		Cfg:              cfg,
		BookingMessaging: k,
	}
}

// Assumetion we have 1 milions requests -> create 1 milions connect to client if the send a message to a topic?
// - The issue occur
// Solution:
// Singleton for create client
