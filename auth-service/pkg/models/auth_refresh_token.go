package models

import (
	"time"
)

type AuthRefreshToken struct {
	Id            int       `gorm:"column:id" json:"id"`
	AuthProfileId int       `gorm:"column:auth_profile_id" json:"auth_profile_id"`
	RefreshToken  string    `gorm:"column:refresh_token" json:"refresh_token"`
	ExpiredAt     time.Time `gorm:"column:expired_at" json:"expired_at"`
	CreatedAt     time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt     time.Time `gorm:"column:updated_at" json:"-"`
}
