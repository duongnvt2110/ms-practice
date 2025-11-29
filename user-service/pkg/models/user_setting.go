package models

import "time"

type UserSetting struct {
	Id        int32     `gorm:"column:id" json:"id"`
	UserId    int32     `gorm:"column:user_id" json:"user_id"`
	AllowNoti bool      `gorm:"column:allow_noti" json:"allow_noti"`
	CreatedAt time.Time `gorm:"column:created_at" json:"-"`
	UpdatedAt time.Time `gorm:"column:updated_at" json:"-"`
}

func (UserSetting) TableName() string {
	return "user_settings"
}
