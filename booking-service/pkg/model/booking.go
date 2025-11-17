package model

import "time"

const (
	BookingStatusPending = "pending"
)

type Booking struct {
	Id             int           `gorm:"primaryKey;column:id" json:"id"`
	IdempotencyKey string        `gorm:"column:idempotency_key;uniqueIndex" json:"idempotency_key"`
	UserId         int           `gorm:"column:user_id" json:"user_id"`
	EventId        int           `gorm:"column:event_id" json:"event_id"`
	BookingCode    string        `gorm:"column:booking_code" json:"booking_code"`
	HoldedAt       time.Time     `gorm:"column:holded_at" json:"holded_at"`
	Status         string        `gorm:"column:status" json:"status"`
	TotalPrice     int           `gorm:"column:total_price" json:"total_price"`
	Logs           string        `gorm:"column:logs" json:"logs"`
	Items          []BookingItem `gorm:"foreignKey:BookingId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"items,omitempty"`
	BookingUsers   []BookingUser `gorm:"foreignKey:BookingId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"booking_users,omitempty"`
	CreatedAt      time.Time     `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time     `gorm:"column:updated_at" json:"updated_at"`
}

func (Booking) TableName() string {
	return "bookings"
}
