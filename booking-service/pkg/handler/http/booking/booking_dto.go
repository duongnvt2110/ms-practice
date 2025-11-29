package booking

type createBookingItemRequest struct {
	TicketTypeID int `json:"ticket_type_id" binding:"required"`
	Qty          int `json:"qty" binding:"required,gt=0"`
	Price        int `json:"price" binding:"required,gte=0"`
}

type createBookingRequest struct {
	UserID      int                        `json:"user_id" binding:"required"`
	EventID     int                        `json:"event_id" binding:"required"`
	BookingCode string                     `json:"booking_code"`
	Status      string                     `json:"status"`
	TotalPrice  int                        `json:"total_price" binding:"omitempty,gte=0"`
	Logs        string                     `json:"logs"`
	Items       []createBookingItemRequest `json:"items" binding:"required"`
}
