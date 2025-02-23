package model

import (
	"time"
)

type AuthProfile struct {
	Id        int       `gorm:"column:id" json:"id"`
	Email     string    `gorm:"column:email" json:"id"`
	Password  string    `gorm:"column:password" json:"id"`
	CreatedAt time.Time `gorm:"column:created_at" json:"id"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"id"`
}
