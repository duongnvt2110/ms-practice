package model

import (
	"time"
)

type Payment struct {
	Id               int              `gorm:"primaryKey;column:id" json:"id"`
	UserId           int              `gorm:"column:user_id" json:"user_id"`
	BookingID        int              `gorm:"column:booking_id" json:"booking_id"`
	Status           string           `gorm:"column:status" json:"status"`
	Quantity         int              `gorm:"column:quantity" json:"quantity"`
	Prices           float64          `gorm:"column:prices" json:"prices"`
	IdempotencyKey   string           `gorm:"column:idempotency_key;uniqueIndex" json:"idempotency_key"`
	PaymentHistories []PaymentHistory `gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"payment_histories,omitempty"`
	CreatedAt        time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time        `gorm:"column:updated_at" json:"updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}
