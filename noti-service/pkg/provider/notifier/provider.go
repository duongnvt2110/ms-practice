package notifier

import (
	"context"
	"log"

	"ms-practice/noti-service/pkg/model"
)

type Provider interface {
	Send(ctx context.Context, job *model.NotificationJob) error
}

type EmailProvider struct {
	From string
}

func (p *EmailProvider) Send(ctx context.Context, job *model.NotificationJob) error {
	log.Printf("sending email notification job=%d template=%s", job.ID, job.Template)
	// TODO: integrate SMTP
	return nil
}

type PushProvider struct{}

func (p *PushProvider) Send(ctx context.Context, job *model.NotificationJob) error {
	log.Printf("sending push notification job=%d", job.ID)
	return nil
}

type InAppProvider struct{}

func (p *InAppProvider) Send(ctx context.Context, job *model.NotificationJob) error {
	log.Printf("recording in-app notification job=%d", job.ID)
	return nil
}
