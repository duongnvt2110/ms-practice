package entity

import "time"

type Booking struct {
	Id        int
	UserId    int
	EventId   int
	Status    string
	Quantity  int
	Prices    float32
	IdemKey   string
	Logs      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
