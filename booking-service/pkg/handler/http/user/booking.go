package user

import (
	"booking-service/pkg/events"
	"booking-service/pkg/utils/response"
	"encoding/json"
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
	ctx := c.Request.Context()
	orderCreated := h.createBooking()
	if orderCreated.OrderID != "" {
		b, _ := json.Marshal(orderCreated)
		_ = h.kafka.SetWriterTopic(events.TopicOrderEvents).Publish(ctx, nil, b)

		paymentProcessed := h.processPayment(orderCreated.OrderID, orderCreated.Amount)

		if paymentProcessed.Success {
			b, _ := json.Marshal(paymentProcessed)
			_ = h.kafka.SetWriterTopic(events.TopicPaymentEvents).Publish(ctx, nil, b)
		} else {
			h.compensateOrder(orderCreated.OrderID, "Payment failed")
		}
	}
	// response.ResponseWithSuccess(c, booking)
}

// Private functions

func (h *bookingHandler) createBooking() events.OrderPlacedEvent {
	return events.OrderPlacedEvent{OrderID: "1", Amount: 1.0}
}

func (h *bookingHandler) processPayment(orderID string, amount float64) events.PaymentProcessedEvent {
	return events.PaymentProcessedEvent{OrderID: "1", PaymentID: "1", Success: true}
}

func (h *bookingHandler) compensateOrder(orderID string, reason string) {
	// Logic to compensate the order, e.g., cancel or mark as failed
	fmt.Println("Compensating Order:", orderID, "Reason:", reason)
}
