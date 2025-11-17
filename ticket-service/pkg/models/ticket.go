package models

import "time"

type Ticket struct {
	Id           int       `gorm:"primaryKey;column:id" json:"id"`
	UserId       int       `gorm:"column:user_id" json:"user_id"`
	BookingId    int       `gorm:"column:booking_id" json:"booking_id"`
	PaymentId    int       `gorm:"column:payment_id" json:"payment_id"`
	TicketTypeId int       `gorm:"column:ticket_type_id" json:"ticket_type_id"`
	Code         string    `gorm:"column:code" json:"code"`
	QRURL        string    `gorm:"column:qr_url" json:"qr_url"`
	Status       string    `gorm:"column:status" json:"status"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (Ticket) TableName() string {
	return "tickets"
}
