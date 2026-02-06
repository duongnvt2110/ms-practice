package model

import (
	"time"
)

type User struct {
	Id           int32         `gorm:"column:id" json:"id"`
	Email        string        `gorm:"column:email" json:"email"`
	Username     string        `gorm:"column:username" json:"username"`
	Avatar       string        `gorm:"column:avatar" json:"avatar"`
	Birthday     string        `gorm:"column:birthday" json:"birthday"`
	MobileNumber string        `gorm:"column:mobile_number" json:"mobile_number"`
	Settings     []UserSetting `gorm:"foreignKey:UserId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"user_settings,omitempty"`
	CreatedAt    time.Time     `gorm:"column:created_at" json:"-"`
	UpdatedAt    time.Time     `gorm:"column:updated_at" json:"-"`
}
