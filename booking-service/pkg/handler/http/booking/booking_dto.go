package booking

type createBookingItemRequest struct {
	EventTypeID int     `json:"event_type_id" binding:"required"`
	Quantity    int     `json:"quantity" binding:"required,gt=0"`
	UnitPrice   float64 `json:"unit_price" binding:"required,gte=0"`
	Currency    string  `json:"currency" binding:"required"`
}

type createBookingRequest struct {
	UserID   int                        `json:"user_id" binding:"required"`
	EventID  int                        `json:"event_id" binding:"required"`
	Quantity int                        `json:"quantity" binding:"required"`
	Prices   float64                    `json:"prices" binding:"required"`
	Items    []createBookingItemRequest `json:"items" binding:"required"`
}
