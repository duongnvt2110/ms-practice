package booking

import (
	"errors"
	"strconv"

	usecase "ms-practice/booking-service/pkg/usecase/booking"
	apperror "ms-practice/booking-service/pkg/util/app_error"
	"ms-practice/booking-service/pkg/util/response"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *bookingHandler) GetBookings(c *gin.Context) {
	ctx := c.Request.Context()
	var userID *int
	if userIDParam := c.Query("user_id"); userIDParam != "" {
		id, err := strconv.Atoi(userIDParam)
		if err != nil {
			response.ResponseWithError(c, apperror.ErrBadRequest.Wrap(err))
			return
		}
		userID = &id
	}

	bookings, err := h.usecase.ListBookings(ctx, usecase.ListBookingsInput{UserID: userID})
	if err != nil {
		response.ResponseWithError(c, apperror.ErrInternalServer.Wrap(err))
		return
	}

	response.ResponseWithSuccess(c, bookings)
}

func (h *bookingHandler) GetBooking(c *gin.Context) {
	ctx := c.Request.Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ResponseWithError(c, apperror.ErrBadRequest.Wrap(err))
		return
	}

	booking, err := h.usecase.GetBooking(ctx, id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			response.ResponseWithError(c, apperror.ErrNotFound.Wrap(err))
			return
		}
		response.ResponseWithError(c, apperror.ErrInternalServer.Wrap(err))
		return
	}

	response.ResponseWithSuccess(c, booking)
}

func (h *bookingHandler) CreateBooking(c *gin.Context) {
	ctx := c.Request.Context()
	var req createBookingRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		response.ResponseWithError(c, apperror.ErrBadRequest.Wrap(err))
		return
	}

	input := usecase.CreateBookingInput{
		UserID:         req.UserID,
		EventID:        req.EventID,
		IdempotencyKey: req.IdempotencyKey,
	}
	if headerKey := c.GetHeader("Idempotency-Key"); headerKey != "" {
		input.IdempotencyKey = headerKey
	}
	for _, item := range req.Items {
		input.Items = append(input.Items, usecase.CreateBookingItem{
			EventTypeID: item.EventTypeID,
			Quantity:    item.Quantity,
			UnitPrice:   item.UnitPrice,
			Currency:    item.Currency,
		})
	}

	result, err := h.usecase.CreateBooking(ctx, input)
	if err != nil {
		var validationErr usecase.ValidationError
		switch {
		case errors.As(err, &validationErr):
			response.ResponseWithError(c, apperror.ErrBadRequest.Wrap(err))
		default:
			response.ResponseWithError(c, apperror.ErrInternalServer.Wrap(err))
		}
		return
	}

	response.ResponseWithSuccess(c, result)
}
