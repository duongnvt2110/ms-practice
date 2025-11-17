package models

import "time"

type TicketType struct {
	Id          int        `gorm:"primaryKey;column:id" json:"id"`
	EventId     int        `gorm:"column:event_id;index" json:"event_id"`
	Position    int        `gorm:"column:position" json:"position"`
	Name        string     `gorm:"column:name" json:"name"`
	Description string     `gorm:"column:description" json:"description"`
	ImageURL    string     `gorm:"column:image_url" json:"image_url"`
	Status      string     `gorm:"column:status" json:"status"`
	Qty         int        `gorm:"column:qty" json:"qty"`
	Price       int        `gorm:"column:price" json:"price"`
	Location    string     `gorm:"column:location" json:"location"`
	SaleAt      *time.Time `gorm:"column:sale_at" json:"sale_at"`
	SaleEnd     *time.Time `gorm:"column:sale_end" json:"sale_end"`
	UpdatedAt   time.Time  `gorm:"column:updated_at" json:"updated_at"`
}

func (TicketType) TableName() string {
	return "ticket_types"
}
