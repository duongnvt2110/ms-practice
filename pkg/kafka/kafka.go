package kafka

import (
	"context"
	"errors"
	"log"
	"ms-practice/booking-service/pkg/event"
	"ms-practice/pkg/backoff"
	"ms-practice/pkg/config"

	"github.com/davecgh/go-spew/spew"
	"github.com/segmentio/kafka-go"
)

const (
	NumberOfWorkers = 3
	NumberOfTasks   = 1000
	MaxRetries      = 3
)

var jobs = make(chan kafka.Message, NumberOfTasks)

type KafkaClient interface {
	Publish(ctx context.Context, key, value []byte) error
	Consume(ctx context.Context, handler func(kafka.Message) error) error
	SetWriterTopic(topic string) KafkaClient
	SetReaderTopic(topic string, groupId string) KafkaClient
}

type kafkaClient struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
	cfg    *config.Kafka
}

func NewKafkaClient(cfg *config.Kafka) KafkaClient {
	return &kafkaClient{
		cfg: cfg,
	}
}

func (k *kafkaClient) SetWriterTopic(topic string) KafkaClient {
	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: k.cfg.Brokers,
		Topic:   topic,
	})
	k.Writer = kWriter
	return k
}

func (k *kafkaClient) SetReaderTopic(topic string, groupId string) KafkaClient {
	kReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: k.cfg.Brokers,
		Topic:   topic,
		GroupID: groupId,
	})
	k.Reader = kReader
	return k
}

func (k *kafkaClient) Publish(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{Key: key, Value: value}
	return k.Writer.WriteMessages(ctx, msg)
}

func (k *kafkaClient) Consume(ctx context.Context, handler func(kafka.Message) error) error {
	k.starWorkerPool(ctx, handler)
	defer close(jobs)
	for {
		msg, err := k.Reader.ReadMessage(ctx)
		if err != nil {
			return err
		}

		if err := handler(msg); err != nil {
			spew.Dump(err)
			select {
			case jobs <- msg:
			case <-ctx.Done():
				return ctx.Err()
			}
		}
	}
}

func (k *kafkaClient) worker(ctx context.Context, workerID int, handler func(kafka.Message) error) {
	b := backoff.NewBackoff(MaxRetries)
	// Define what errors are retryable (customize this).
	isRetryable := func(err error) bool {
		// Example: retry everything for now.
		// Skip Retry
		// In production, exclude poison pill / validation errors.
		if errors.Is(err, context.Canceled) ||
			errors.Is(err, context.DeadlineExceeded) {
			return false
		}

		return true
	}
	for retryMsg := range jobs {
		err := b.Retry(ctx, b, isRetryable, func() error {
			log.Printf("Worker %d: handling message %s\n", workerID, string(retryMsg.Value))
			return handler(retryMsg)
		})
		if err != nil {
			// Retries exhausted OR non-retryable error
			log.Printf("Worker %d: Move to DLQ message %s (err: %v)\n",
				workerID, string(retryMsg.Value), err)
			k.SetWriterTopic(event.DLQTopicName)
			k.Publish(ctx, nil, retryMsg.Value)
		}
	}
}

func (k *kafkaClient) starWorkerPool(ctx context.Context, handler func(kafka.Message) error) {
	for i := 0; i < NumberOfWorkers; i++ {
		go k.worker(ctx, i, handler)
	}
}
