package event

import (
	"ms-practice/event-service/pkg/config"
	"ms-practice/event-service/pkg/usecase"
	eventUC "ms-practice/event-service/pkg/usecase/event"
)

type Handler struct {
	cfg     *config.Config
	eventUC eventUC.Usecase
}

func NewHandler(cfg *config.Config, uc *usecase.Usecase) Handler {
	return Handler{
		cfg:     cfg,
		eventUC: uc.EventUC,
	}
}
