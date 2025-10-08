package model

import (
	"time"
)

type Booking struct {
	Id             int           `gorm:"primaryKey;column:id" json:"id"`
	UserId         int           `gorm:"column:user_id" json:"user_id"`
	EventId        int           `gorm:"column:event_id" json:"event_id"`
	Status         string        `gorm:"column:status" json:"status"`
	Quantity       int           `gorm:"column:quantity" json:"quantity"`
	Prices         float64       `gorm:"column:prices" json:"prices"`
	IdempotencyKey string        `gorm:"column:idempotency_key;uniqueIndex" json:"idempotency_key"`
	Logs           string        `gorm:"column:logs" json:"logs"`
	Items          []BookingItem `gorm:"foreignKey:BookingId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"items,omitempty"`
	CreatedAt      time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

func (Booking) TableName() string {
	return "bookings"
}
