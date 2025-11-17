package model

import (
	"time"
)

type Payment struct {
	Id               int              `gorm:"primaryKey;column:id" json:"id"`
	IdempotencyKey   string           `gorm:"column:idempotency_key" json:"idempotency_key"`
	UserId           int              `gorm:"column:user_id" json:"user_id"`
	BookingID        int              `gorm:"column:booking_id" json:"booking_id"`
	TransactionID    string           `gorm:"column:transaction_id" json:"transaction_id"`
	PaymentCode      string           `gorm:"column:payment_code" json:"payment_code"`
	Status           string           `gorm:"column:status" json:"status"`
	Provider         string           `gorm:"column:provider" json:"provider"`
	Prices           int              `gorm:"column:prices" json:"prices"`
	PaidAt           time.Time        `gorm:"column:paid_at" json:"paid_at"`
	PaymentHistories []PaymentHistory `gorm:"foreignKey:PaymentID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"payment_histories,omitempty"`
	CreatedAt        time.Time        `gorm:"column:created_at" json:"created_at"`
	UpdatedAt        time.Time        `gorm:"column:updated_at" json:"updated_at"`
}

func (Payment) TableName() string {
	return "payments"
}
