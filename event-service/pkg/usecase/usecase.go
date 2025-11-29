package usecase

import (
	"ms-practice/event-service/pkg/repositories"
	"ms-practice/event-service/pkg/usecase/event"
)

type Usecase struct {
	EventUC event.Usecase
}

func NewUsecase(repo *repositories.Repository) *Usecase {
	return &Usecase{
		EventUC: event.NewUsecase(repo.EventRepo),
	}
}
