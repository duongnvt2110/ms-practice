package repositories

import (
	ticketRepo "ms-practice/ticket-service/pkg/repositories/ticket"

	"gorm.io/gorm"
)

type Repository struct {
	TicketRepo ticketRepo.Repository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		TicketRepo: ticketRepo.NewRepository(db),
	}
}
