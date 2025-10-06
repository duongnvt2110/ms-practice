package entity

import "time"

type BookingItem struct {
	Id          int
	BookingId   int
	EventTypeId int
	Qty         string
	UnitPrice   int
	Currency    string
	CreatedAt   time.Time
	UpdatedAt   time.Time
}
