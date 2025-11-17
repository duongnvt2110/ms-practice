package ticket

import (
	"ms-practice/ticket-service/pkg/config"
	"ms-practice/ticket-service/pkg/usecase"
	ticketUC "ms-practice/ticket-service/pkg/usecase/ticket"
)

type Handler struct {
	cfg      *config.Config
	ticketUC ticketUC.Usecase
}

func NewHandler(cfg *config.Config, uc *usecase.Usecase) Handler {
	return Handler{
		cfg:      cfg,
		ticketUC: uc.TicketUC,
	}
}
