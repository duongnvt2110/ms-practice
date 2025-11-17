package booking

import (
	"context"
	"ms-practice/booking-service/pkg/model"

	"gorm.io/gorm"
)

type BookingRepository interface {
	Create(ctx context.Context, booking *model.Booking) error
	GetByID(ctx context.Context, id int) (*model.Booking, error)
	GetByIdempotencyKey(ctx context.Context, key string) (*model.Booking, error)
	List(ctx context.Context, userID *int) ([]model.Booking, error)
}

type bookingRepository struct {
	db *gorm.DB
}

func NewBookingRepository(db *gorm.DB) BookingRepository {
	return &bookingRepository{
		db: db,
	}
}

func (r *bookingRepository) Create(ctx context.Context, booking *model.Booking) error {
	return r.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(booking).Error; err != nil {
			return err
		}
		return nil
	})
}

func (r *bookingRepository) GetByID(ctx context.Context, id int) (*model.Booking, error) {
	var booking model.Booking
	err := r.db.WithContext(ctx).
		Preload("Items").
		Preload("BookingUsers").
		First(&booking, id).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) GetByIdempotencyKey(ctx context.Context, key string) (*model.Booking, error) {
	var booking model.Booking
	err := r.db.WithContext(ctx).
		Preload("Items").
		Preload("BookingUsers").
		Where("idempotency_key = ?", key).
		First(&booking).Error
	if err != nil {
		return nil, err
	}
	return &booking, nil
}

func (r *bookingRepository) List(ctx context.Context, userID *int) ([]model.Booking, error) {
	var bookings []model.Booking
	query := r.db.WithContext(ctx).
		Model(&model.Booking{}).
		Preload("Items").
		Preload("BookingUsers").
		Order("created_at DESC")
	if userID != nil {
		query = query.Where("user_id = ?", *userID)
	}
	err := query.Find(&bookings).Error
	if err != nil {
		return nil, err
	}
	return bookings, nil
}
