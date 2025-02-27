package repositories

import "gorm.io/gorm"

type Repository struct {
	AuthProfileRepo AuthProfileRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{AuthProfileRepo: NewAuthProfileRepo(db)}
}
