package model

import "time"

type DLQRecord struct {
	ID          uint64    `gorm:"primaryKey;autoIncrement"`
	Topic       string    `gorm:"size:255;not null"`
	Partition   int       `gorm:"column:partition_id;not null"`
	Offset      int64     `gorm:"column:offset_id;not null"`
	Key         []byte    `gorm:"column:key;type:blob"`
	Headers     string    `gorm:"type:longtext"`
	Payload     []byte    `gorm:"type:longblob;not null"`
	PayloadJSON *string   `gorm:"column:payload_json;type:longtext"`
	CreatedAt   time.Time `gorm:"autoCreateTime"`
}

func (DLQRecord) TableName() string {
	return "dlq_records"
}
