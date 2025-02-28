package repositories

import "gorm.io/gorm"

type Repository struct {
	AuthProfileRepo  AuthProfileRepo
	RefreshTokenRepo RefreshTokenRepo
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		AuthProfileRepo:  NewAuthProfileRepo(db),
		RefreshTokenRepo: NewRefreshTokenRepo(db),
	}
}
