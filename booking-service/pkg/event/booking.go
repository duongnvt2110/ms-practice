package events

// OrderPlacedEvent represents an order creation message sent by OrderService
// during the choreography saga.
type BookingOrdered struct {
	EventType string  `json:"event_type"`
	OrderID   string  `json:"order_id"`
	Amount    float64 `json:"amount"`
}
