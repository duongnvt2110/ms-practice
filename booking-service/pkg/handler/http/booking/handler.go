package booking

import (
	"ms-practice/booking-service/pkg/config"
	usecase "ms-practice/booking-service/pkg/usecase/booking"
)

type bookingHandler struct {
	cfg     *config.Config
	usecase usecase.Usecase
}

func NewBookingHandler(cfg *config.Config, uc usecase.Usecase) *bookingHandler {
	return &bookingHandler{
		cfg:     cfg,
		usecase: uc,
	}
}
