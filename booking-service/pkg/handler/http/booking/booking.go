package booking

import (
	"errors"
	"ms-practice/booking-service/pkg/model"
	apperror "ms-practice/booking-service/pkg/util/app_error"
	"ms-practice/booking-service/pkg/util/response"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (h *bookingHandler) GetBookings(c *gin.Context) {
	ctx := c.Request.Context()

	var userID int
	if userIDParam := c.Query("user_id"); userIDParam != "" {
		id, err := strconv.Atoi(userIDParam)
		if err != nil {
			response.ResponseWithError(c, apperror.ErrBadRequest.Wrap(err))
			return
		}
		userID = id
	}

	bookings, err := h.bookingUC.ListBookings(ctx, userID)
	if err != nil {
		response.ResponseWithError(c, apperror.ErrInternalServer.Wrap(err))
		return
	}

	response.ResponseWithSuccess(c, bookings)
}

func (h *bookingHandler) GetBooking(c *gin.Context) {
	// Retrieve the `id` from the route
	ctx := c.Request.Context()
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		response.ResponseWithError(c, apperror.ErrBadRequest.Wrap(err))
		return
	}

	booking, err := h.bookingUC.GetBooking(ctx, id)
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

	booking := &model.Booking{
		EventId:  req.EventID,
		Quantity: req.Quantity,
		Prices:   req.Prices,
	}

	if headerKey := c.GetHeader("Idempotency-Key"); headerKey != "" {
		booking.IdempotencyKey = headerKey
	}

	for _, item := range req.Items {
		booking.Items = append(booking.Items, model.BookingItem{
			EventTypeId: item.EventTypeID,
			Qty:         item.Quantity,
			UnitPrice:   item.UnitPrice,
			Currency:    item.Currency,
		})
	}
	result, err := h.bookingUC.CreateBooking(ctx, booking)
	if err != nil {
		response.ResponseWithError(c, apperror.ErrInternalServer.Wrap(err))
		return
	}
	response.ResponseWithSuccess(c, result.Id)
}

// Private functions

// func (h *bookingHandler) processPayment(orderID string, amount float64) events.PaymentProcessedEvent {
// 	return events.PaymentProcessedEvent{OrderID: "1", PaymentID: "1", Success: true}
// }

// func (h *bookingHandler) compensateOrder(orderID string, reason string) {
// 	// Logic to compensate the order, e.g., cancel or mark as failed
// 	fmt.Println("Compensating Order:", orderID, "Reason:", reason)
// }
