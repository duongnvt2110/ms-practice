package event

import (
	"context"
	"ms-practice/catalog-service/pkg/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, event *models.Event) error
	GetByID(ctx context.Context, id int) (*models.Event, error)
	List(ctx context.Context, status string) ([]models.Event, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, event *models.Event) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(event).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repository) GetByID(ctx context.Context, id int) (*models.Event, error) {
	var event models.Event
	if err := r.db.WithContext(ctx).
		Preload("TicketTypes").
		First(&event, id).Error; err != nil {
		return nil, err
	}
	return &event, nil
}

func (r *repository) List(ctx context.Context, status string) ([]models.Event, error) {
	var events []models.Event
	query := r.db.WithContext(ctx).Model(&models.Event{}).Preload("TicketTypes").Order("start_at ASC")
	if status != "" {
		query = query.Where("status = ?", status)
	}
	if err := query.Find(&events).Error; err != nil {
		return nil, err
	}
	return events, nil
}
