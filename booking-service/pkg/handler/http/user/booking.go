package user

import (
	"encoding/json"
	"fmt"
	"ms-practice/booking-service/pkg/events"
	"ms-practice/booking-service/pkg/utils/response"

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
		Name: "demo n times",
	}
	response.ResponseWithSuccess(c, booking)
}

func (h *bookingHandler) CreateBooking(c *gin.Context) {
	ctx := c.Request.Context()
	orderCreated := h.createBooking()
	if orderCreated.OrderID != "" {
		b, _ := json.Marshal(orderCreated)
		err := h.bookingMessaging.Producers[events.BookingTopicName].Publish(ctx, nil, b)
		fmt.Println(err)
	}
	response.ResponseWithSuccess(c, "test")
}

// Private functions

func (h *bookingHandler) createBooking() events.BookingOrdered {
	return events.BookingOrdered{
		EventType: "BookingOrdered",
		OrderID:   "2",
		Amount:    1.0,
	}
}

// func (h *bookingHandler) processPayment(orderID string, amount float64) events.PaymentProcessedEvent {
// 	return events.PaymentProcessedEvent{OrderID: "1", PaymentID: "1", Success: true}
// }

// func (h *bookingHandler) compensateOrder(orderID string, reason string) {
// 	// Logic to compensate the order, e.g., cancel or mark as failed
// 	fmt.Println("Compensating Order:", orderID, "Reason:", reason)
// }
