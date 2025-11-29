package event

import (
	"context"
	"errors"
	"ms-practice/event-service/pkg/models"
	eventRepo "ms-practice/event-service/pkg/repositories/event"
	"time"
)

const (
	defaultEventStatus = "draft"
)

type Usecase interface {
	CreateEvent(ctx context.Context, event *models.Event) error
	GetEvent(ctx context.Context, id int) (*models.Event, error)
	ListEvents(ctx context.Context, status string) ([]models.Event, error)
}

type usecase struct {
	eventRepo eventRepo.Repository
}

func NewUsecase(eventRepo eventRepo.Repository) Usecase {
	return &usecase{eventRepo: eventRepo}
}

func (u *usecase) CreateEvent(ctx context.Context, event *models.Event) error {
	if event == nil {
		return errors.New("event payload is required")
	}
	setEventDefaults(event)
	return u.eventRepo.Create(ctx, event)
}

func (u *usecase) GetEvent(ctx context.Context, id int) (*models.Event, error) {
	return u.eventRepo.GetByID(ctx, id)
}

func (u *usecase) ListEvents(ctx context.Context, status string) ([]models.Event, error) {
	return u.eventRepo.List(ctx, status)
}

func setEventDefaults(event *models.Event) {
	if event.Status == "" {
		event.Status = defaultEventStatus
	}
	if event.StartAt.IsZero() {
		event.StartAt = time.Now()
	}
	if event.EndAt.IsZero() {
		event.EndAt = event.StartAt.Add(3 * time.Hour)
	}
}
