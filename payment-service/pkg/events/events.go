package events

// OrderPlacedEvent represents an order creation message sent by OrderService
// during the choreography saga.
type OrderPlacedEvent struct {
	OrderID string  `json:"order_id"`
	Amount  float64 `json:"amount"`
}

// PaymentProcessedEvent represents the outcome of the PaymentService
// after an order has been paid for.
type PaymentProcessedEvent struct {
	OrderID   string `json:"order_id"`
	PaymentID string `json:"payment_id"`
	Success   bool   `json:"success"`
}

// CompensateOrderEvent is emitted when a previous step fails and a
// rollback is required.
type CompensateOrderEvent struct {
	OrderID string `json:"order_id"`
	Reason  string `json:"reason"`
}

const (
	TopicOrderEvents   = "order-events"
	TopicPaymentEvents = "payment-events"
)
