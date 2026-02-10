package event

type TicketPayload struct {
	EventType  TicketEventType `json:"event_type"`
	TicketID   int             `json:"ticket_id"`
	BookingID  int             `json:"booking_id"`
	UserID     int             `json:"user_id"`
	Email      string          `json:"email,omitempty"`
	DeviceIDs  []string        `json:"device_ids,omitempty"`
	TicketCode string          `json:"ticket_code,omitempty"`
}

type TicketEventType string

const (
	TicketIssued TicketEventType = "TicketIssued"
)
