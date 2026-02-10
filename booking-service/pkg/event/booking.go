package event

type BookingPayload struct {
	EventType BookingEventType `json:"event_type"`
	OrderID   int              `json:"order_id"`
	UserID    int              `json:"user_id"`
	Amount    float64          `json:"amount"`
}

type BookingEventType string

const (
	BookingOrdered   BookingEventType = "BookingOrdered"
	BookingCancelled BookingEventType = "BookingCancelled"
)
