package kafka

import (
	"context"
	"log"
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
	k.starWorkerPool(handler)
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
			default:
				log.Print("Move to DLQ")
			}
		}

	}
}

func (k *kafkaClient) worker(workerID int, handler func(kafka.Message) error) {
	for retryMsg := range jobs {
		retries := 1
		for {
			log.Printf("Worker %d Retry %v message %v\n", workerID, retries, string(retryMsg.Value))
			if retries >= MaxRetries {
				// Move to DLQ
				log.Printf("Move to DLQ\n")
				break
			}
			if err := handler(retryMsg); err != nil {
				log.Printf("Error processing message, retrying: %v\n", err)
				retries++
				continue
			}
			break
		}
	}
}

func (k *kafkaClient) starWorkerPool(handler func(kafka.Message) error) {
	for i := 0; i < NumberOfWorkers; i++ {
		go k.worker(i, handler)
	}
}
