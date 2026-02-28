package consumer

import (
	"context"
	"errors"
	"log"
	"time"

	"ms-practice/retry-management-service/pkg/config"
	"ms-practice/retry-management-service/pkg/usecase"

	"github.com/segmentio/kafka-go"
)

func StartDLQConsumer(ctx context.Context, cfg *config.Config, dlqUC usecase.DLQUsecase) error {
	if cfg == nil || cfg.App.Kafka == nil {
		return errors.New("kafka config is not set")
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:        cfg.App.Kafka.Brokers,
		Topic:          cfg.DlqTopic,
		GroupID:        cfg.DlqGroupID,
		MinBytes:       10e3,
		MaxBytes:       10e6,
		MaxWait:        500 * time.Millisecond,
		CommitInterval: 0,
	})
	defer func() {
		if err := reader.Close(); err != nil {
			log.Printf("failed to close kafka reader: %v", err)
		}
	}()

	log.Printf("DLQ consumer started for topic=%s group=%s", cfg.DlqTopic, cfg.DlqGroupID)

	for {
		msg, err := reader.FetchMessage(ctx)
		if err != nil {
			if errors.Is(err, context.Canceled) || errors.Is(err, context.DeadlineExceeded) {
				return nil
			}
			log.Printf("kafka fetch error: %v", err)
			continue
		}

		if err := dlqUC.Ingest(ctx, msg); err != nil {
			log.Printf("dlq ingest error: %v", err)
			continue
		}

		if err := reader.CommitMessages(ctx, msg); err != nil {
			log.Printf("kafka commit error: %v", err)
		}
	}
}
