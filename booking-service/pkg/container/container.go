package container

import (
	"booking-service/pkg/config"
	"ms-practice/pkg/kafka"
)

type Container struct {
	Cfg   *config.Config
	Kafka kafka.KafkaClient
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	k := kafka.NewKafkaClient(cfg.App.Kafka)
	return &Container{
		Cfg:   cfg,
		Kafka: k,
	}
}
