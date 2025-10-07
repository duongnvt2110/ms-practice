package booking

import (
	"ms-practice/booking-service/pkg/config"
	"ms-practice/booking-service/pkg/util/kafka"
)

type bookingHandler struct {
	cfg              *config.Config
	bookingMessaging *kafka.BookingMessaging
}

func NewBookingHandler(cfg *config.Config, BookingMessaging *kafka.BookingMessaging) bookingHandler {
	return bookingHandler{
		cfg:              cfg,
		bookingMessaging: BookingMessaging,
	}
}
