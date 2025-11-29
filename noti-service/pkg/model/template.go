package model

import "time"

type NotificationTemplate struct {
    ID        uint64 `gorm:"primaryKey;column:id"`
    Code      string `gorm:"column:code;uniqueIndex"`
    Channel   NotificationChannel `gorm:"column:channel"`
    Version   int    `gorm:"column:version"`
    Subject   string `gorm:"column:subject"`
    Body      string `gorm:"column:body"`
    CreatedAt time.Time
    UpdatedAt time.Time
}

func (NotificationTemplate) TableName() string {
    return "notification_templates"
}
