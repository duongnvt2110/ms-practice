package event

type Topic struct {
	Name           string
	GroupID        string
	EnableProducer bool
	EnableConsumer bool
}

const (
	BookingTopicName = "booking.events"
	PaymentTopicName = "payments.events"
	TicketTopicName  = "tickets.events"
	DLQTopicName     = "dlq.events"
)
