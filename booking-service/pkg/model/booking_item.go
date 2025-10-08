package model

import "time"

type BookingItem struct {
	Id          int       `gorm:"primaryKey;column:id" json:"id"`
	BookingId   int       `gorm:"column:booking_id;index" json:"booking_id"`
	EventTypeId int       `gorm:"column:event_type_id" json:"event_type_id"`
	Qty         int       `gorm:"column:qty" json:"quantity"`
	UnitPrice   float64   `gorm:"column:unit_price" json:"unit_price"`
	Currency    string    `gorm:"column:currency" json:"currency"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BookingItem) TableName() string {
	return "booking_items"
}
