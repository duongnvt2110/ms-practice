package e2e

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strconv"
	"testing"

	"ms-practice/booking-service/pkg/config"
	bookingHTTP "ms-practice/booking-service/pkg/handler/http/booking"
	"ms-practice/booking-service/pkg/model"
	bookingRepo "ms-practice/booking-service/pkg/repository/booking"
	bookingUsecase "ms-practice/booking-service/pkg/usecase/booking"
	"ms-practice/booking-service/pkg/util/response"

	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/stretchr/testify/require"
	"gorm.io/gorm"
)

type recorderPublisher struct {
	published []*model.Booking
}

func (r *recorderPublisher) BookingCreated(ctx context.Context, booking *model.Booking) error {
	r.published = append(r.published, booking)
	return nil
}

func setupE2EDB(t *testing.T) *gorm.DB {
	t.Helper()
	db, err := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	require.NoError(t, err)
	require.NoError(t, db.AutoMigrate(&model.Booking{}, &model.BookingItem{}))
	return db
}

func TestBookingWorkflow_EndToEnd(t *testing.T) {
	gin.SetMode(gin.TestMode)

	db := setupE2EDB(t)
	repo := bookingRepo.NewRepository(db)
	publisher := &recorderPublisher{}
	usecase := bookingUsecase.NewUsecase(repo, publisher)
	cfg := &config.Config{}
	handler := bookingHTTP.NewBookingHandler(cfg, usecase)

	router := gin.New()
	router.POST("/bookings", handler.CreateBooking)
	router.GET("/bookings/:id", handler.GetBooking)
	router.GET("/bookings", handler.GetBookings)

	payload := map[string]interface{}{
		"user_id":         1,
		"event_id":        99,
		"idempotency_key": "order-1",
		"items": []map[string]interface{}{
			{"event_type_id": 1001, "quantity": 2, "unit_price": 15000, "currency": "VND"},
			{"event_type_id": 1002, "quantity": 1, "unit_price": 20000, "currency": "VND"},
		},
	}
	body, _ := json.Marshal(payload)
	createReq := httptest.NewRequest(http.MethodPost, "/bookings", bytes.NewBuffer(body))
	createReq.Header.Set("Content-Type", "application/json")
	createRec := httptest.NewRecorder()

	router.ServeHTTP(createRec, createReq)
	require.Equal(t, http.StatusOK, createRec.Code)
	var createResp response.APIResponse
	require.NoError(t, json.Unmarshal(createRec.Body.Bytes(), &createResp))

	dataBytes, err := json.Marshal(createResp.Data)
	require.NoError(t, err)
	var created model.Booking
	require.NoError(t, json.Unmarshal(dataBytes, &created))
	require.NotZero(t, created.Id)
	require.Equal(t, 3, created.Quantity)
	require.Len(t, publisher.published, 1)

	getReq := httptest.NewRequest(http.MethodGet, "/bookings/"+strconv.Itoa(created.Id), nil)
	getRec := httptest.NewRecorder()
	router.ServeHTTP(getRec, getReq)
	require.Equal(t, http.StatusOK, getRec.Code)

	listReq := httptest.NewRequest(http.MethodGet, "/bookings?user_id=1", nil)
	listRec := httptest.NewRecorder()
	router.ServeHTTP(listRec, listReq)
	require.Equal(t, http.StatusOK, listRec.Code)
}
