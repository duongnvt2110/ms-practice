package booking

import (
	"context"

	"ms-practice/booking-service/pkg/model"

	"gorm.io/gorm"
)

type ListInput struct {
	UserID *int
}

type Repository interface {
	Create(ctx context.Context, booking *model.Booking) error
	GetByID(ctx context.Context, id int) (*model.Booking, error)
	GetByIdempotencyKey(ctx context.Context, key string) (*model.Booking, error)
	List(ctx context.Context, input ListInput) ([]model.Booking, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) Repository {
	return &repository{db: db}
}

func (r *repository) Create(ctx context.Context, booking *model.Booking) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(booking).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *repository) GetByID(ctx context.Context, id int) (*model.Booking, error) {
	var booking model.Booking
	err := r.db.WithContext(ctx).Preload("Items").First(&booking, id).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *repository) GetByIdempotencyKey(ctx context.Context, key string) (*model.Booking, error) {
	var booking model.Booking
	err := r.db.WithContext(ctx).Preload("Items").Where("idempotency_key = ?", key).First(&booking).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *repository) List(ctx context.Context, input ListInput) ([]model.Booking, error) {
	var bookings []model.Booking
	query := r.db.WithContext(ctx).Model(&model.Booking{}).Preload("Items").Order("created_at DESC")
	if input.UserID != nil {
		query = query.Where("user_id = ?", *input.UserID)
	}
	err := query.Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
