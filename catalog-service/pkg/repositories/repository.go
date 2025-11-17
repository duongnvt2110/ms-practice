package repositories

import (
	"ms-practice/catalog-service/pkg/repositories/event"

	"gorm.io/gorm"
)

type Repository struct {
	EventRepo event.Repository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		EventRepo: event.NewRepository(db),
	}
}
