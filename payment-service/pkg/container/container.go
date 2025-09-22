package container

import (
	"payment-service/pkg/config"
	kafka_client "payment-service/pkg/utils/kafka"
)

type Container struct {
	Cfg   *config.Config
	Kafka kafka_client.KafkaClient
}

func InitializeContainer() *Container {
	cfg := config.NewConfig()
	kafka := kafka_client.NewKafkaClient(cfg)
	return &Container{Cfg: cfg, Kafka: kafka}
}
