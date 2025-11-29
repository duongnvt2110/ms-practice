package model

import "time"

type BookingItem struct {
	Id           int       `gorm:"primaryKey;column:id" json:"id"`
	BookingId    int       `gorm:"column:booking_id;index" json:"booking_id"`
	TicketTypeId int       `gorm:"column:ticket_type_id" json:"ticket_type_id"`
	Qty          int       `gorm:"column:qty" json:"qty"`
	Price        int       `gorm:"column:price" json:"price"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BookingItem) TableName() string {
	return "booking_items"
}
