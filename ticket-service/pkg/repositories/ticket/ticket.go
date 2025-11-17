package ticket

import (
	"context"
	"ms-practice/ticket-service/pkg/models"

	"gorm.io/gorm"
)

type Repository interface {
	Create(ctx context.Context, ticket *models.Ticket) error
	GetByID(ctx context.Context, id int) (*models.Ticket, error)
	List(ctx context.Context, userID *int) ([]models.Ticket, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, ticket *models.Ticket) error {
	return r.db.WithContext(ctx).Create(ticket).Error
}

func (r *repository) GetByID(ctx context.Context, id int) (*models.Ticket, error) {
	var ticket models.Ticket
	if err := r.db.WithContext(ctx).First(&ticket, id).Error; err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *repository) List(ctx context.Context, userID *int) ([]models.Ticket, error) {
	var tickets []models.Ticket
	query := r.db.WithContext(ctx).Model(&models.Ticket{}).Order("created_at DESC")
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	if err := query.Find(&tickets).Error; err != nil {
		return nil, err
	}
	return tickets, nil
}
