package repository

import (
	"context"
	"time"

	"ms-practice/noti-service/pkg/model"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type NotificationRepository interface {
	CreateJobs(ctx context.Context, jobs []model.NotificationJob) error
	FindPendingJob(ctx context.Context, channel model.NotificationChannel) (*model.NotificationJob, error)
	UpdateJob(ctx context.Context, job *model.NotificationJob) error
}

type notificationRepository struct {
	db *gorm.DB
}

func NewNotificationRepository(db *gorm.DB) NotificationRepository {
	return &notificationRepository{db: db}
}

func (r *notificationRepository) CreateJobs(ctx context.Context, jobs []model.NotificationJob) error {
	if len(jobs) == 0 {
		return nil
	}
	now := time.Now()
	for i := range jobs {
		if jobs[i].Status == "" {
			jobs[i].Status = model.StatusPending
		}
		if jobs[i].NextAttemptAt.IsZero() {
			jobs[i].NextAttemptAt = now
		}
	}
	return r.db.WithContext(ctx).Create(&jobs).Error
}

func (r *notificationRepository) FindPendingJob(ctx context.Context, channel model.NotificationChannel) (*model.NotificationJob, error) {
	var job model.NotificationJob
	err := r.db.WithContext(ctx).
		Clauses(clause.Locking{Strength: "UPDATE", Options: "SKIP LOCKED"}).
		Where("channel = ? AND status = ? AND next_attempt_at <= ?", channel, model.StatusPending, time.Now()).
		Order("id").
		First(&job).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, err
	}
	return &job, nil
}

func (r *notificationRepository) UpdateJob(ctx context.Context, job *model.NotificationJob) error {
	return r.db.WithContext(ctx).Save(job).Error
}
