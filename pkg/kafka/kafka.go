package kafka

import (
	"context"
	"ms-practice/pkg/config"

	"github.com/segmentio/kafka-go"
)

type KafkaClient interface {
	Publish(ctx context.Context, key, value []byte) error
	Consume(ctx context.Context, handler func(kafka.Message)) error
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

func (k *kafkaClient) Consume(ctx context.Context, handler func(kafka.Message)) error {
	for {
		m, err := k.Reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		handler(m)
	}
}
