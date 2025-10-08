package usecase

import "ms-practice/payment-service/pkg/repository"

type Usecase struct {
	repo *repository.Repository
}

func NewUsecase(repo *repository.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}
}
