package event

import (
	"ms-practice/catalog-service/pkg/config"
	"ms-practice/catalog-service/pkg/usecase"
	eventUC "ms-practice/catalog-service/pkg/usecase/event"
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
