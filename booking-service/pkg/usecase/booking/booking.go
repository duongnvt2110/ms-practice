package booking

import (
	"context"
	"encoding/json"
	"errors"
	"ms-practice/booking-service/pkg/event"
	events "ms-practice/booking-service/pkg/event"
	"ms-practice/booking-service/pkg/model"
	"ms-practice/booking-service/pkg/repository/booking"
	"ms-practice/booking-service/pkg/util/kafka"

	"gorm.io/gorm"
)

type BookingUsecase interface {
	CreateBooking(ctx context.Context, booking *model.Booking) (*model.Booking, error)
	GetBooking(ctx context.Context, id int) (*model.Booking, error)
	ListBookings(ctx context.Context, userID int) ([]model.Booking, error)
}

type bookingUsecase struct {
	bookingRepo booking.BookingRepository
	messaging   *kafka.BookingMessaging
}

func NewBookingUsecase(bookingRepo booking.BookingRepository, messaging *kafka.BookingMessaging) BookingUsecase {
	return &bookingUsecase{
		bookingRepo: bookingRepo,
		messaging:   messaging,
	}
}

func (u *bookingUsecase) CreateBooking(ctx context.Context, booking *model.Booking) (*model.Booking, error) {
	if booking.IdempotencyKey != "" {
		if existing, err := u.bookingRepo.GetByIdempotencyKey(ctx, booking.IdempotencyKey); err == nil {
			return existing, nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	if err := u.bookingRepo.Create(ctx, booking); err != nil {
		return nil, err
	}

	if err := u.BookingCreated(ctx, booking); err != nil {
		return nil, err
	}

	return booking, nil
}

func (u *bookingUsecase) GetBooking(ctx context.Context, id int) (*model.Booking, error) {
	return u.bookingRepo.GetByID(ctx, id)
}

func (u *bookingUsecase) ListBookings(ctx context.Context, userID int) ([]model.Booking, error) {
	return u.bookingRepo.List(ctx, &userID)
}

func (u *bookingUsecase) BookingCreated(ctx context.Context, booking *model.Booking) error {
	if u.messaging == nil {
		return nil
	}
	producer, ok := u.messaging.Producers[event.BookingTopicName]
	if !ok || producer == nil {
		return nil
	}

	payload := events.BookingPayload{
		EventType: string(event.BookingOrdered),
		OrderID:   booking.Id,
		Amount:    booking.Prices,
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return producer.Publish(ctx, nil, data)
}
