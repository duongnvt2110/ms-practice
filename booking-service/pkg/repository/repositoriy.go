package repository

import (
	"ms-practice/booking-service/pkg/repository/booking"

	"gorm.io/gorm"
)

type Repository struct {
	BookingRepo booking.BookingRepository
}

func NewRepository(db *gorm.DB) *Repository {
	bookingUC := booking.NewBookingRepository(db)
	return &Repository{
		BookingRepo: bookingUC,
	}
}
