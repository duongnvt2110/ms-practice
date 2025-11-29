package model

import "time"

type BookingUser struct {
	Id           int       `gorm:"primaryKey;column:id" json:"id"`
	BookingId    int       `gorm:"column:booking_id;index" json:"booking_id"`
	UserId       int       `gorm:"column:user_id" json:"user_id"`
	Email        string    `gorm:"column:email" json:"email"`
	MobileNumber string    `gorm:"column:mobile_number" json:"mobile_number"`
	Address      string    `gorm:"column:address" json:"address"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"created_at"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"updated_at"`
}

func (BookingUser) TableName() string {
	return "booking_users"
}
