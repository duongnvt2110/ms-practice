package worker

import (
	"context"
	"log"
	"time"

	"ms-practice/noti-service/pkg/config"
	"ms-practice/noti-service/pkg/model"
	"ms-practice/noti-service/pkg/provider/notifier"
	"ms-practice/noti-service/pkg/repository"
)

type Worker struct {
	Channel  model.NotificationChannel
	Repo     repository.NotificationRepository
	Provider notifier.Provider
	Cfg      *config.Config
}

func (w *Worker) Start(ctx context.Context) {
	if w.Provider == nil || w.Repo == nil {
		log.Printf("worker for channel %s missing dependencies", w.Channel)
		return
	}
	interval := time.Duration(w.Cfg.Notification.WorkerIntervalMS) * time.Millisecond
	ticker := time.NewTicker(interval)
	defer ticker.Stop()
	for {
		select {
		case <-ctx.Done():
			log.Printf("worker for %s stopping", w.Channel)
			return
		case <-ticker.C:
			if err := w.process(ctx); err != nil {
				log.Printf("worker %s process error: %v", w.Channel, err)
			}
		}
	}
}

func (w *Worker) process(ctx context.Context) error {
	job, err := w.Repo.FindPendingJob(ctx, w.Channel)
	if err != nil {
		return err
	}
	if job == nil {
		return nil
	}
	job.Status = model.StatusSending
	job.Attempts++
	if err := w.Repo.UpdateJob(ctx, job); err != nil {
		return err
	}

	sendErr := w.Provider.Send(ctx, job)
	if sendErr != nil {
		job.Status = model.StatusPending
		job.LastError = sendErr.Error()
		job.NextAttemptAt = time.Now().Add(time.Duration(w.Cfg.Notification.RetryIntervalSeconds) * time.Second)
		if job.Attempts >= job.MaxAttempts {
			job.Status = model.StatusFailed
		}
	} else {
		now := time.Now()
		job.Status = model.StatusSent
		job.SentAt = &now
		job.LastError = ""
	}
	return w.Repo.UpdateJob(ctx, job)
}
