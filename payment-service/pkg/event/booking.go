package event

// OrderPlacedEvent represents an order creation message sent by OrderService
// during the choreography saga.
type BookingPayload struct {
	EventType string  `json:"event_type"`
	OrderID   int     `json:"order_id"`
	Amount    float64 `json:"amount"`
}

type BookingEventType string

const (
	BookingOrdered   BookingEventType = "BookingOrdered"
	BookingCancelled BookingEventType = "BookingCancelled"
)
