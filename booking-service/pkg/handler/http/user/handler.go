package user

import (
	"booking-service/pkg/config"
	"ms-practice/pkg/kafka"
)

type bookingHandler struct {
	cfg   *config.Config
	kafka kafka.KafkaClient
}

func NewBookingHandler(cfg *config.Config, kafka kafka.KafkaClient) bookingHandler {
	return bookingHandler{
		cfg:   cfg,
		kafka: kafka,
	}
}
