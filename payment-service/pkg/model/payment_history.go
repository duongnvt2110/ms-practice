package model

import (
	"time"
)

type PaymentHistory struct {
	Id        int       `gorm:"primaryKey;column:id" json:"id"`
	PaymentID int       `gorm:"column:payment_id" json:"payment_id"`
	Status    int       `gorm:"column:status" json:"status"`
	Logs      string    `gorm:"column:logs" json:"logs"`
	PaidAt    int       `gorm:"column:paid_at" json:"paid_at"`
	CreatedAt time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (PaymentHistory) TableName() string {
	return "payment_histories"
}
