package repository

import (
	"context"
	"errors"
	"fmt"

	"ms-practice/retry-management-service/pkg/model"

	"gorm.io/gorm"
)

type DLQRepository interface {
	Create(ctx context.Context, record *model.DLQRecord) error
	List(ctx context.Context, page, pageSize int) ([]model.DLQRecord, int64, error)
	GetByID(ctx context.Context, id int64) (*model.DLQRecord, error)
}

type dlqRepository struct {
	db *gorm.DB
}

func NewDLQRepository(db *gorm.DB) DLQRepository {
	return &dlqRepository{db: db}
}

func (r *dlqRepository) Create(ctx context.Context, record *model.DLQRecord) error {
	if record == nil {
		return fmt.Errorf("dlq record is nil")
	}
	return r.db.WithContext(ctx).Create(record).Error
}

func (r *dlqRepository) List(ctx context.Context, page, pageSize int) ([]model.DLQRecord, int64, error) {
	if page <= 0 {
		page = 1
	}
	if pageSize <= 0 || pageSize > 200 {
		pageSize = 20
	}

	var total int64
	if err := r.db.WithContext(ctx).Model(&model.DLQRecord{}).Count(&total).Error; err != nil {
		return nil, 0, err
	}

	var records []model.DLQRecord
	err := r.db.WithContext(ctx).
		Order("id DESC").
		Limit(pageSize).
		Offset((page - 1) * pageSize).
		Find(&records).Error
	if err != nil {
		return nil, 0, err
	}
	return records, total, nil
}

func (r *dlqRepository) GetByID(ctx context.Context, id int64) (*model.DLQRecord, error) {
	if id <= 0 {
		return nil, fmt.Errorf("invalid id")
	}
	var record model.DLQRecord
	err := r.db.WithContext(ctx).First(&record, "id = ?", id).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &record, nil
}
