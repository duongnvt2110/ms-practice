package event

type PaymentPayload struct {
	EventType PaymentEventType `json:"event_type"`
	OrderID   int              `json:"order_id"`
	UserID    int              `json:"user_id"`
	PaymentID int              `json:"payment_id"`
	Amount    float64          `json:"amount"`
	Message   string           `json:"message,omitempty"`
	Email     string           `json:"email,omitempty"`
	DeviceIDs []string         `json:"device_ids,omitempty"`
}

type PaymentEventType string

const (
	PaymentSucceeded PaymentEventType = "PaymentSucceeded"
	PaymentFailed    PaymentEventType = "PaymentFailed"
)
