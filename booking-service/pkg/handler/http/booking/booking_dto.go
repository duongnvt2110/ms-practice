package booking

type BookingFormRequest struct {
	EventID     int `json:"event_id" binding:"required"`
	EventItemId int `json:"event_item_id" binding:"required"`
	Quantity    int `json:"quantity" binding:"required"`
	// IdempotencyKey string `json:"idempotency_key" binding:"required"`
}
