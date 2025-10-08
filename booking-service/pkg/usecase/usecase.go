package usecase

import (
	"ms-practice/booking-service/pkg/repository"
	"ms-practice/booking-service/pkg/usecase/booking"
	"ms-practice/booking-service/pkg/util/kafka"
)

type Usecase struct {
	BookingUC booking.BookingUsecase
}

func NewUsecase(repo *repository.Repository, message *kafka.BookingMessaging) *Usecase {
	bookingUC := booking.NewBookingUsecase(repo.BookingRepo, message)
	return &Usecase{
		BookingUC: bookingUC,
	}
}
