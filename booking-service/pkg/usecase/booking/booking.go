package booking

import (
	"context"
	"crypto/rand"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"ms-practice/booking-service/pkg/event"
	events "ms-practice/booking-service/pkg/event"
	"ms-practice/booking-service/pkg/model"
	"ms-practice/booking-service/pkg/repository/booking"
	"ms-practice/booking-service/pkg/util/kafka"
	"strings"
	"time"

	"gorm.io/gorm"
)

type BookingUsecase interface {
	CreateBooking(ctx context.Context, booking *model.Booking) (*model.Booking, error)
	GetBooking(ctx context.Context, id int) (*model.Booking, error)
	ListBookings(ctx context.Context, userID *int) ([]model.Booking, error)
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

	setBookingDefaults(booking)

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

func (u *bookingUsecase) ListBookings(ctx context.Context, userID *int) ([]model.Booking, error) {
	return u.bookingRepo.List(ctx, userID)
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
		Amount:    float64(booking.TotalPrice),
	}

	data, err := json.Marshal(payload)
	if err != nil {
		return err
	}

	return producer.Publish(ctx, nil, data)
}

func setBookingDefaults(booking *model.Booking) {
	if booking.IdempotencyKey == "" {
		booking.IdempotencyKey = randomHexString(32)
	}

	if booking.BookingCode == "" {
		booking.BookingCode = fmt.Sprintf("BK-%s", strings.ToUpper(randomHexString(8)))
	}

	if booking.Status == "" {
		booking.Status = model.BookingStatusPending
	}

	if booking.Logs == "" {
		booking.Logs = "{}"
	}

	if booking.TotalPrice == 0 {
		booking.TotalPrice = calculateTotalPrice(booking.Items)
	}
}

func calculateTotalPrice(items []model.BookingItem) int {
	total := 0
	for _, item := range items {
		total += item.Price * item.Qty
	}
	return total
}

func randomHexString(length int) string {
	if length <= 0 {
		return ""
	}
	byteLen := (length + 1) / 2
	buf := make([]byte, byteLen)
	if _, err := rand.Read(buf); err != nil {
		return fmt.Sprintf("%d", time.Now().UnixNano())
	}
	encoded := hex.EncodeToString(buf)
	if len(encoded) >= length {
		return encoded[:length]
	}
	return encoded
}
