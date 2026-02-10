package repository

import "gorm.io/gorm"

type Repository struct {
	UserRepo UserRepository
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		UserRepo: NewUserRepository(db),
	}
}
