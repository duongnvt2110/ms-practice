package booking_test

import (
	"context"
	"errors"
	"testing"

	"ms-practice/booking-service/pkg/model"
	repo "ms-practice/booking-service/pkg/repository/booking"
	booking "ms-practice/booking-service/pkg/usecase/booking"

	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type stubRepo struct {
	existingByKey *model.Booking
	createInput   *model.Booking
	createErr     error
}

func (s *stubRepo) Create(ctx context.Context, b *model.Booking) error {
	s.createInput = b
	return s.createErr
}

func (s *stubRepo) GetByIdempotencyKey(ctx context.Context, key string) (*model.Booking, error) {
	if s.existingByKey != nil {
		return s.existingByKey, nil
	}
	return nil, gorm.ErrRecordNotFound
}

func (s *stubRepo) GetByID(ctx context.Context, id int) (*model.Booking, error) {
	return nil, errors.New("not implemented")
}

func (s *stubRepo) List(ctx context.Context, input repo.ListInput) ([]model.Booking, error) {
	return nil, nil
}

type stubPublisher struct {
	called  bool
	payload *model.Booking
	err     error
}

func (s *stubPublisher) BookingCreated(ctx context.Context, booking *model.Booking) error {
	s.called = true
	s.payload = booking
	if s.err != nil {
		return s.err
	}
	return nil
}

func TestCreateBooking_InvalidInput(t *testing.T) {
	repo := &stubRepo{}
	publisher := &stubPublisher{}
	usecase := booking.NewUsecase(repo, publisher)

	_, err := usecase.CreateBooking(context.Background(), booking.CreateBookingInput{})
	require.Error(t, err)
	require.False(t, publisher.called)
	require.Nil(t, repo.createInput)
}

func TestCreateBooking_Idempotent(t *testing.T) {
	existing := &model.Booking{Id: 42, IdempotencyKey: "dup-key"}
	repo := &stubRepo{existingByKey: existing}
	publisher := &stubPublisher{}
	usecase := booking.NewUsecase(repo, publisher)

	result, err := usecase.CreateBooking(context.Background(), booking.CreateBookingInput{
		UserID:         1,
		EventID:        2,
		IdempotencyKey: "dup-key",
		Items: []booking.CreateBookingItem{
			{EventTypeID: 3, Quantity: 1, UnitPrice: 1000, Currency: "VND"},
		},
	})
	require.NoError(t, err)
	require.Equal(t, existing.Id, result.Id)
	require.Nil(t, repo.createInput)
	require.False(t, publisher.called)
}

func TestCreateBooking_Success(t *testing.T) {
	repo := &stubRepo{}
	publisher := &stubPublisher{}
	usecase := booking.NewUsecase(repo, publisher)

	result, err := usecase.CreateBooking(context.Background(), booking.CreateBookingInput{
		UserID:         11,
		EventID:        22,
		IdempotencyKey: "unique-key",
		Items: []booking.CreateBookingItem{
			{EventTypeID: 101, Quantity: 2, UnitPrice: 15000, Currency: "VND"},
			{EventTypeID: 102, Quantity: 1, UnitPrice: 20000, Currency: "VND"},
		},
	})
	require.NoError(t, err)

	require.NotNil(t, repo.createInput)
	require.Equal(t, 3, repo.createInput.Quantity)
	require.InDelta(t, 50000, repo.createInput.Prices, 0.01)
	require.Equal(t, "pending", repo.createInput.Status)
	require.Len(t, repo.createInput.Items, 2)

	require.True(t, publisher.called)
	require.NotNil(t, publisher.payload)
	require.Equal(t, result, publisher.payload)
}

func TestCreateBooking_PropagatesRepositoryError(t *testing.T) {
	expectedErr := errors.New("db failure")
	repo := &stubRepo{createErr: expectedErr}
	publisher := &stubPublisher{}
	usecase := booking.NewUsecase(repo, publisher)

	_, err := usecase.CreateBooking(context.Background(), booking.CreateBookingInput{
		UserID:         1,
		EventID:        2,
		IdempotencyKey: "key",
		Items: []booking.CreateBookingItem{
			{EventTypeID: 3, Quantity: 1, UnitPrice: 1000, Currency: "VND"},
		},
	})
	require.ErrorIs(t, err, expectedErr)
	require.False(t, publisher.called)
}

func TestCreateBooking_PublisherError(t *testing.T) {
	expectedErr := errors.New("publish failed")
	repo := &stubRepo{}
	publisher := &stubPublisher{err: expectedErr}
	usecase := booking.NewUsecase(repo, publisher)

	_, err := usecase.CreateBooking(context.Background(), booking.CreateBookingInput{
		UserID:         12,
		EventID:        34,
		IdempotencyKey: "key",
		Items: []booking.CreateBookingItem{
			{EventTypeID: 3, Quantity: 1, UnitPrice: 1200, Currency: "USD"},
		},
	})
	require.ErrorIs(t, err, expectedErr)
	require.True(t, publisher.called)
}
