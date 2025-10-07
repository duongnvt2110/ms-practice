package booking

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"ms-practice/booking-service/pkg/config"
	"ms-practice/booking-service/pkg/model"
	usecase "ms-practice/booking-service/pkg/usecase/booking"
	"ms-practice/booking-service/pkg/util/response"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/require"
)

type stubUsecase struct {
	createFn func(context.Context, usecase.CreateBookingInput) (*model.Booking, error)
	getFn    func(context.Context, int) (*model.Booking, error)
	listFn   func(context.Context, usecase.ListBookingsInput) ([]model.Booking, error)
}

func (s *stubUsecase) CreateBooking(ctx context.Context, input usecase.CreateBookingInput) (*model.Booking, error) {
	return s.createFn(ctx, input)
}

func (s *stubUsecase) GetBooking(ctx context.Context, id int) (*model.Booking, error) {
	return s.getFn(ctx, id)
}

func (s *stubUsecase) ListBookings(ctx context.Context, input usecase.ListBookingsInput) ([]model.Booking, error) {
	return s.listFn(ctx, input)
}

func setupRouter(handler *bookingHandler) *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.New()
	router.POST("/bookings", handler.CreateBooking)
	router.GET("/bookings/:id", handler.GetBooking)
	router.GET("/bookings", handler.GetBookings)
	return router
}

func TestCreateBookingHandler_Success(t *testing.T) {
	usecase := &stubUsecase{
		createFn: func(ctx context.Context, input usecase.CreateBookingInput) (*model.Booking, error) {
			return &model.Booking{
				Id:             1,
				UserId:         input.UserID,
				EventId:        input.EventID,
				Status:         "pending",
				Quantity:       2,
				Prices:         20000,
				IdempotencyKey: input.IdempotencyKey,
			}, nil
		},
		getFn: func(context.Context, int) (*model.Booking, error) { return nil, nil },
		listFn: func(context.Context, usecase.ListBookingsInput) ([]model.Booking, error) {
			return nil, nil
		},
	}

	handler := NewBookingHandler(&config.Config{}, usecase)
	router := setupRouter(handler)

	payload := map[string]interface{}{
		"user_id":         10,
		"event_id":        20,
		"idempotency_key": "key",
		"items": []map[string]interface{}{
			{"event_type_id": 30, "quantity": 2, "unit_price": 10000, "currency": "VND"},
		},
	}

	body, _ := json.Marshal(payload)
	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)

	var resp response.APIResponse
	require.NoError(t, json.Unmarshal(rec.Body.Bytes(), &resp))
	require.Equal(t, http.StatusOK, resp.Status)
	require.NotNil(t, resp.Data)
}

func TestCreateBookingHandler_ValidationError(t *testing.T) {
	usecase := &stubUsecase{
		createFn: func(ctx context.Context, input usecase.CreateBookingInput) (*model.Booking, error) {
			t.Fatalf("usecase should not be called")
			return nil, nil
		},
		getFn:  func(context.Context, int) (*model.Booking, error) { return nil, nil },
		listFn: func(context.Context, usecase.ListBookingsInput) ([]model.Booking, error) { return nil, nil },
	}

	handler := NewBookingHandler(&config.Config{}, usecase)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBufferString(`{}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusBadRequest, rec.Code)
}

func TestCreateBookingHandler_UsecaseError(t *testing.T) {
	usecase := &stubUsecase{
		createFn: func(ctx context.Context, input usecase.CreateBookingInput) (*model.Booking, error) {
			return nil, errors.New("boom")
		},
		getFn:  func(context.Context, int) (*model.Booking, error) { return nil, nil },
		listFn: func(context.Context, usecase.ListBookingsInput) ([]model.Booking, error) { return nil, nil },
	}

	handler := NewBookingHandler(&config.Config{}, usecase)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBufferString(`{
		"user_id": 1,
		"event_id": 2,
		"idempotency_key": "k",
		"items": [{"event_type_id": 3, "quantity": 1, "unit_price": 1000, "currency": "VND"}]
	}`))
	req.Header.Set("Content-Type", "application/json")
	rec := httptest.NewRecorder()

	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusInternalServerError, rec.Code)
}

func TestGetBookingHandler_Success(t *testing.T) {
	usecase := &stubUsecase{
		createFn: func(context.Context, usecase.CreateBookingInput) (*model.Booking, error) { return nil, nil },
		getFn: func(ctx context.Context, id int) (*model.Booking, error) {
			return &model.Booking{Id: id, UserId: 10, EventId: 20}, nil
		},
		listFn: func(context.Context, usecase.ListBookingsInput) ([]model.Booking, error) { return nil, nil },
	}
	handler := NewBookingHandler(&config.Config{}, usecase)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/bookings/5", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}

func TestGetBookingsHandler_Success(t *testing.T) {
	usecase := &stubUsecase{
		createFn: func(context.Context, usecase.CreateBookingInput) (*model.Booking, error) { return nil, nil },
		getFn:    func(context.Context, int) (*model.Booking, error) { return nil, nil },
		listFn: func(ctx context.Context, input usecase.ListBookingsInput) ([]model.Booking, error) {
			return []model.Booking{
				{Id: 1, UserId: 1, EventId: 100},
				{Id: 2, UserId: 1, EventId: 101},
			}, nil
		},
	}
	handler := NewBookingHandler(&config.Config{}, usecase)
	router := setupRouter(handler)

	req := httptest.NewRequest(http.MethodGet, "/bookings?user_id=1", nil)
	rec := httptest.NewRecorder()
	router.ServeHTTP(rec, req)

	require.Equal(t, http.StatusOK, rec.Code)
}
