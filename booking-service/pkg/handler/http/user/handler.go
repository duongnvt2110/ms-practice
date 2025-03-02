package user

import "booking-service/pkg/config"


type bookingHandler struct {
	cfg *config.Config
}

func NewBookingHandler(cfg *config.Config) bookingHandler {
	return bookingHandler{
		cfg: cfg,
	}
}
