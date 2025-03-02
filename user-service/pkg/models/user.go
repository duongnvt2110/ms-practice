package models

import (
	"time"
)

type User struct {
	Id          int32     `gorm:"column:id" json:"id"`
	Email       string    `gorm:"column:email" json:"email"`
	FirstName   string    `gorm:"column:first_name" json:"first_name"`
	LastName    string    `gorm:"column:last_name" json:"last_name"`
	Birthday    string    `gorm:"column:birthday" json:"birthday"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	CreatedAt   time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt   time.Time `gorm:"column:updated_at" json:"-"`
}
