package model

import (
	"time"
)

type NotificationChannel string

type NotificationStatus string

const (
	ChannelEmail NotificationChannel = "email"
	ChannelPush  NotificationChannel = "push"
	ChannelInApp NotificationChannel = "in_app"

	StatusPending NotificationStatus = "pending"
	StatusSending NotificationStatus = "sending"
	StatusSent    NotificationStatus = "sent"
	StatusFailed  NotificationStatus = "failed"
)

// NotificationJob represents a per-channel delivery task.
type NotificationJob struct {
	ID             uint64              `gorm:"primaryKey;column:id" json:"id"`
	IdempotencyKey string              `gorm:"column:idempotency_key;uniqueIndex" json:"idempotency_key"`
	EventType      string              `gorm:"column:event_type" json:"event_type"`
	EventID        int                 `gorm:"column:event_id" json:"event_id"`
	UserID         int                 `gorm:"column:user_id" json:"user_id"`
	Channel        NotificationChannel `gorm:"column:channel" json:"channel"`
	Template       string              `gorm:"column:template" json:"template"`
	Payload        string              `gorm:"column:payload" json:"payload"`
	Status         NotificationStatus  `gorm:"column:status" json:"status"`
	Attempts       int                 `gorm:"column:attempts" json:"attempts"`
	MaxAttempts    int                 `gorm:"column:max_attempts" json:"max_attempts"`
	LastError      string              `gorm:"column:last_error" json:"last_error"`
	NextAttemptAt  time.Time           `gorm:"column:next_attempt_at" json:"next_attempt_at"`
	SentAt         *time.Time          `gorm:"column:sent_at" json:"sent_at"`
	CreatedAt      time.Time           `gorm:"column:created_at" json:"created_at"`
	UpdatedAt      time.Time           `gorm:"column:updated_at" json:"updated_at"`
}

func (NotificationJob) TableName() string {
	return "notification_jobs"
}
