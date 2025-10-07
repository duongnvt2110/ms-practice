package booking

import (
	"context"
	"errors"
	"fmt"

	"ms-practice/booking-service/pkg/model"
	repo "ms-practice/booking-service/pkg/repository/booking"

	"gorm.io/gorm"
)

type BookingPublisher interface {
	BookingCreated(ctx context.Context, booking *model.Booking) error
}

type CreateBookingItem struct {
	EventTypeID int
	Quantity    int
	UnitPrice   float64
	Currency    string
}

type CreateBookingInput struct {
	UserID         int
	EventID        int
	IdempotencyKey string
	Items          []CreateBookingItem
}

type ListBookingsInput struct {
	UserID *int
}

type Usecase interface {
	CreateBooking(ctx context.Context, input CreateBookingInput) (*model.Booking, error)
	GetBooking(ctx context.Context, id int) (*model.Booking, error)
	ListBookings(ctx context.Context, input ListBookingsInput) ([]model.Booking, error)
}

type usecase struct {
	repo      repo.Repository
	publisher BookingPublisher
}

func NewUsecase(repo repo.Repository, publisher BookingPublisher) Usecase {
	return &usecase{
		repo:      repo,
		publisher: publisher,
	}
}

func (u *usecase) CreateBooking(ctx context.Context, input CreateBookingInput) (*model.Booking, error) {
	if err := validateCreateInput(input); err != nil {
		return nil, err
	}

	if input.IdempotencyKey != "" {
		if existing, err := u.repo.GetByIdempotencyKey(ctx, input.IdempotencyKey); err == nil {
			return existing, nil
		} else if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, err
		}
	}

	entity := buildBookingEntity(input)
	if err := u.repo.Create(ctx, entity); err != nil {
		return nil, err
	}

	if err := u.publisher.BookingCreated(ctx, entity); err != nil {
		return nil, err
	}

	return entity, nil
}

func (u *usecase) GetBooking(ctx context.Context, id int) (*model.Booking, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *usecase) ListBookings(ctx context.Context, input ListBookingsInput) ([]model.Booking, error) {
	repoInput := repo.ListInput{
		UserID: input.UserID,
	}
	return u.repo.List(ctx, repoInput)
}

func validateCreateInput(input CreateBookingInput) error {
	if input.UserID <= 0 {
		return ValidationError{message: "user_id must be greater than zero"}
	}
	if input.EventID <= 0 {
		return ValidationError{message: "event_id must be greater than zero"}
	}
	if len(input.Items) == 0 {
		return ValidationError{message: "at least one item is required"}
	}
	for i, item := range input.Items {
		if item.EventTypeID <= 0 {
			return ValidationError{message: fmt.Sprintf("items[%d].event_type_id must be greater than zero", i)}
		}
		if item.Quantity <= 0 {
			return ValidationError{message: fmt.Sprintf("items[%d].quantity must be greater than zero", i)}
		}
		if item.UnitPrice < 0 {
			return ValidationError{message: fmt.Sprintf("items[%d].unit_price must be non-negative", i)}
		}
		if item.Currency == "" {
			return ValidationError{message: fmt.Sprintf("items[%d].currency is required", i)}
		}
	}
	return nil
}

func buildBookingEntity(input CreateBookingInput) *model.Booking {
	items := make([]model.BookingItem, 0, len(input.Items))
	totalQuantity := 0
	totalAmount := float64(0)
	for _, item := range input.Items {
		totalQuantity += item.Quantity
		totalAmount += item.UnitPrice * float64(item.Quantity)
		items = append(items, model.BookingItem{
			EventTypeId: item.EventTypeID,
			Qty:         item.Quantity,
			UnitPrice:   item.UnitPrice,
			Currency:    item.Currency,
		})
	}

	return &model.Booking{
		UserId:         input.UserID,
		EventId:        input.EventID,
		Status:         model.BookingStatusPending,
		Quantity:       totalQuantity,
		Prices:         totalAmount,
		IdempotencyKey: input.IdempotencyKey,
		Logs:           "{}",
		Items:          items,
	}
}

type ValidationError struct {
	message string
}

func (e ValidationError) Error() string {
	return e.message
}
