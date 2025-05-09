package user

import (
	"booking-service/pkg/utils/response"
	"fmt"

	"github.com/gin-gonic/gin"
)

func (h *bookingHandler) GetBookings(c *gin.Context) {
	bookings := []struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		{
			ID:   "123",
			Name: "Test",
		},
		{
			ID:   "123232",
			Name: "Test3",
		},
	}
	response.ResponseWithSuccess(c, bookings)
}

func (h *bookingHandler) GetBooking(c *gin.Context) {
	// Retrieve the `id` from the route
	booking_id := c.Param("id")
	booking := struct {
		ID   string `json:"id"`
		Name string `json:"name"`
	}{
		ID:   booking_id,
		Name: "Test",
	}
	response.ResponseWithSuccess(c, booking)
}

func (h *bookingHandler) CreateBooking(c *gin.Context) {
	orderCreated := h.createBooking()
	if orderCreated.OrderID != "" {
		paymentProcessed := h.processPayment(orderCreated.OrderID, orderCreated.Amount)

		if !paymentProcessed.PaymentDone {
			// Compensate for the order creation due to payment failure
			// This could involve canceling the order or marking it as failed
			h.compensateOrder(orderCreated.OrderID, "Payment failed")
		}
	}
	// response.ResponseWithSuccess(c, booking)
}

// Private func
type OrderCreatedEvent struct {
	OrderID string
	Amount  float64
}

type PaymentProcessedEvent struct {
	OrderID     string
	PaymentID   string
	PaymentDone bool
}

type CompensateOrderEvent struct {
	OrderID string
	Reason  string
}

func (h *bookingHandler) createBooking() OrderCreatedEvent {
	return OrderCreatedEvent{OrderID: "1", Amount: 1.0}
}

func (h *bookingHandler) processPayment(orderID string, amount float64) PaymentProcessedEvent {
	return PaymentProcessedEvent{OrderID: "1", PaymentID: "1", PaymentDone: true}
}

func (h *bookingHandler) compensateOrder(orderID string, reason string) {
	// Logic to compensate the order, e.g., cancel or mark as failed
	fmt.Println("Compensating Order:", orderID, "Reason:", reason)
}
