package models

import "time"

type Event struct {
	Id          int          `gorm:"primaryKey;column:id" json:"id"`
	Name        string       `gorm:"column:name" json:"name"`
	Title       string       `gorm:"column:title" json:"title"`
	Banner      string       `gorm:"column:banner" json:"banner"`
	StartAt     time.Time    `gorm:"column:start_at" json:"start_at"`
	EndAt       time.Time    `gorm:"column:end_at" json:"end_at"`
	Location    string       `gorm:"column:location" json:"location"`
	Status      string       `gorm:"column:status" json:"status"`
	CreatedAt   time.Time    `gorm:"column:created_at" json:"created_at"`
	UpdatedAt   time.Time    `gorm:"column:updated_at" json:"updated_at"`
	TicketTypes []TicketType `gorm:"foreignKey:EventId;constraint:OnUpdate:CASCADE,OnDelete:CASCADE" json:"ticket_types,omitempty"`
}

func (Event) TableName() string {
	return "events"
}
