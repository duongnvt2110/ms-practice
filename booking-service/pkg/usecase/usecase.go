package usecase

import (
	repository "ms-practice/booking-service/pkg/repository"
	bookingUsecase "ms-practice/booking-service/pkg/usecase/booking"
)

type Usecase struct {
	Booking bookingUsecase.Usecase
}

func NewUsecase(repo *repository.Repository, bookingPublisher bookingUsecase.BookingPublisher) *Usecase {
	return &Usecase{
		Booking: bookingUsecase.NewUsecase(repo.Booking, bookingPublisher),
	}
}
