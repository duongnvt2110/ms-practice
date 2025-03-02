package models

import (
	"time"
)

type RefreshToken struct {
	Id        int       `gorm:"column:id" json:"id"`
	UserId    int       `gorm:"column:user_id" json:"user_id"`
	Token     string    `gorm:"column:token" json:"token"`
	ExpiresAt time.Time `gorm:"column:expires_at" json:"expires_at"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}
