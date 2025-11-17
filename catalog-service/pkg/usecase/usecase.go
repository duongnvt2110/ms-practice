package usecase

import (
	"ms-practice/catalog-service/pkg/repositories"
	"ms-practice/catalog-service/pkg/usecase/event"
)

type Usecase struct {
	EventUC event.Usecase
}

func NewUsecase(repo *repositories.Repository) *Usecase {
	return &Usecase{
		EventUC: event.NewUsecase(repo.EventRepo),
	}
}
