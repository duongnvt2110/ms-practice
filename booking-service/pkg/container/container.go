package container

import (
	"booking-service/pkg/config"
	"booking-service/pkg/kafka"
)

type Container struct {
	Cfg   *config.Config
	Kafka kafka.KafkaClient
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	k := kafka.NewKafkaClient(cfg)
	return &Container{
		Cfg:   cfg,
		Kafka: k,
	}
}
