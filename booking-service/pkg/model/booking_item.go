package model

import "time"

type BookingItem struct {
	Id          int
	BookingId   int
	EventTypeId int
	Qty         int
	UnitPrice   int
	Currency    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
