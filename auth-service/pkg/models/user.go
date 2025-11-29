package models

import (
	"time"
)

type User struct {
	Id           string    `gorm:"column:id" json:"id"`
	Email        string    `gorm:"column:email" json:"email"`
	Username     string    `gorm:"column:username" json:"username"`
	Avatar       string    `gorm:"column:avatar" json:"avatar"`
	Birthday     string    `gorm:"column:birthday" json:"birthday"`
	MobileNumber string    `gorm:"column:mobile_number" json:"mobile_number"`
	CreatedAt    time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time `gorm:"column:updated_at" json:"-"`
}
