package repository

import (
	bookingRepo "ms-practice/booking-service/pkg/repository/booking"

	"gorm.io/gorm"
)

type Repository struct {
	Booking bookingRepo.Repository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Booking: bookingRepo.NewRepository(db),
	}
}
