package booking_test

import (
	"context"
	"testing"

	"ms-practice/booking-service/pkg/model"
	"ms-practice/booking-service/pkg/repository/booking"

	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

func setupTestDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Booking{}, &model.BookingItem{}))
	return db
}

func TestCreateAndFetchBooking(t *testing.T) {
	db := setupTestDB(t)
	repo := booking.NewRepository(db)
	ctx := context.Background()

	entity := &model.Booking{
		UserId:         10,
		EventId:        22,
		Status:         model.BookingStatusPending,
		Quantity:       2,
		Prices:         50000,
		IdempotencyKey: "unique-key",
		Logs:           "{}",
		Items: []model.BookingItem{
			{
				EventTypeId: 31,
				Qty:         2,
				UnitPrice:   25000,
				Currency:    "VND",
			},
		},
	}

	err := repo.Create(ctx, entity)
	require.NoError(t, err)
	require.NotZero(t, entity.Id)
	require.Len(t, entity.Items, 1)
	require.NotZero(t, entity.Items[0].Id)

	found, err := repo.GetByID(ctx, entity.Id)
	require.NoError(t, err)
	require.Equal(t, entity.Id, found.Id)
	require.Len(t, found.Items, 1)
	require.Equal(t, 31, found.Items[0].EventTypeId)

	byKey, err := repo.GetByIdempotencyKey(ctx, entity.IdempotencyKey)
	require.NoError(t, err)
	require.Equal(t, entity.Id, byKey.Id)
}
