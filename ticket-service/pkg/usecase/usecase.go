package usecase

import (
	"ms-practice/ticket-service/pkg/repositories"
	ticketUC "ms-practice/ticket-service/pkg/usecase/ticket"
)

type Usecase struct {
	TicketUC ticketUC.Usecase
}

func NewUsecase(repo *repositories.Repository) *Usecase {
	return &Usecase{
		TicketUC: ticketUC.NewUsecase(repo.TicketRepo),
	}
}
