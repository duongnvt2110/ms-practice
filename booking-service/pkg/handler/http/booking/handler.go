package booking

import (
	"ms-practice/booking-service/pkg/config"
	"ms-practice/booking-service/pkg/usecase/booking"
)

type bookingHandler struct {
	cfg       *config.Config
	bookingUC booking.BookingUsecase
}

func NewBookingHandler(cfg *config.Config, bookingUC booking.BookingUsecase) bookingHandler {
	return bookingHandler{
		cfg:       cfg,
		bookingUC: bookingUC,
	}
}
