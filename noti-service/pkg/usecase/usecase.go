package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"ms-practice/noti-service/pkg/config"
	"ms-practice/noti-service/pkg/model"
	"ms-practice/noti-service/pkg/repository"
	"ms-practice/pkg/event"
)

type NotificationUsecase interface {
	HandlePaymentEvent(ctx context.Context, payload event.PaymentPayload) error
	HandleTicketEvent(ctx context.Context, payload event.TicketPayload) error
}

type notificationUsecase struct {
	repo repository.NotificationRepository
	cfg  *config.Config
}

func NewNotificationUsecase(repo repository.NotificationRepository, cfg *config.Config) NotificationUsecase {
	return &notificationUsecase{repo: repo, cfg: cfg}
}

func (u *notificationUsecase) HandlePaymentEvent(ctx context.Context, payload event.PaymentPayload) error {
	jobs := u.buildJobs(ctx, string(payload.EventType), payload.OrderID, payload.UserID, map[string]any{
		"payment_id": payload.PaymentID,
		"amount":     payload.Amount,
		"message":    payload.Message,
	})
	return u.repo.CreateJobs(ctx, jobs)
}

func (u *notificationUsecase) HandleTicketEvent(ctx context.Context, payload event.TicketPayload) error {
	jobs := u.buildJobs(ctx, string(payload.EventType), payload.TicketID, payload.UserID, map[string]any{
		"booking_id": payload.BookingID,
		"ticket_id":  payload.TicketID,
	})
	return u.repo.CreateJobs(ctx, jobs)
}

func (u *notificationUsecase) buildJobs(ctx context.Context, eventType string, eventID int, userID int, payload map[string]any) []model.NotificationJob {
	data, _ := json.Marshal(payload)
	now := time.Now()
	channels := []model.NotificationChannel{model.ChannelEmail, model.ChannelPush, model.ChannelInApp}
	jobs := make([]model.NotificationJob, 0, len(channels))
	for _, ch := range channels {
		jobs = append(jobs, model.NotificationJob{
			IdempotencyKey: fmt.Sprintf("%s-%d-%s", eventType, eventID, ch),
			EventType:      eventType,
			EventID:        eventID,
			UserID:         userID,
			Channel:        ch,
			Template:       templateFor(eventType, ch),
			Payload:        string(data),
			Status:         model.StatusPending,
			Attempts:       0,
			MaxAttempts:    u.cfg.Notification.MaxAttempts,
			NextAttemptAt:  now,
		})
	}
	return jobs
}

func templateFor(eventType string, channel model.NotificationChannel) string {
	switch eventType {
	case string(event.PaymentSucceeded):
		switch channel {
		case model.ChannelEmail:
			return "payment_receipt_email_v1"
		case model.ChannelPush:
			return "payment_receipt_push_v1"
		case model.ChannelInApp:
			return "payment_receipt_inapp_v1"
		}
	case string(event.TicketIssued):
		switch channel {
		case model.ChannelEmail:
			return "ticket_issued_email_v1"
		case model.ChannelPush:
			return "ticket_issued_push_v1"
		case model.ChannelInApp:
			return "ticket_issued_inapp_v1"
		}
	}
	return fmt.Sprintf("%s_%s", eventType, channel)
}
