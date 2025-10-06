package container

import (
	"ms-practice/payment-service/pkg/config"
	kafka_client "ms-practice/payment-service/pkg/util/kafka"
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
