package kafka_client

import (
	"context"

	"payment-service/pkg/config"

	"github.com/segmentio/kafka-go"
)

type KafkaClient interface {
	// Publish sends a message to the configured topic
	Publish(ctx context.Context, key, value []byte) error
	// Consume reads messages from the configured topic and calls handler for
	// each message. The consumer stops when ctx is cancelled or an error occurs.
	Consume(ctx context.Context, handler func(kafka.Message)) error
	// SetWriterTopic sets the topic for the writer
	SetWriterTopic(topic string) KafkaClient
	// SetReaderTopic sets the topic for the reader along with consumer group id
	SetReaderTopic(topic string, groupId string) KafkaClient
}

type kafkaClient struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
	cfg    *config.Config
}

func NewKafkaClient(cfg *config.Config) KafkaClient {
	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: cfg.Kafka.Brokers,
	})
	return &kafkaClient{
		Writer: kWriter,
		cfg:    cfg,
	}
}

func (k *kafkaClient) SetWriterTopic(topic string) KafkaClient {
	k.Writer.Topic = topic
	return k
}

func (k *kafkaClient) SetReaderTopic(topic string, groupId string) KafkaClient {
	kReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: k.cfg.Kafka.Brokers,
		Topic:   topic,
		GroupID: groupId,
	})
	k.Reader = kReader
	return k
}

// Publish implements KafkaClient.Publish
func (k *kafkaClient) Publish(ctx context.Context, key, value []byte) error {
	msg := kafka.Message{Key: key, Value: value}
	return k.Writer.WriteMessages(ctx, msg)
}

// Consume implements KafkaClient.Consume
func (k *kafkaClient) Consume(ctx context.Context, handler func(kafka.Message)) error {
	for {
		m, err := k.Reader.ReadMessage(ctx)
		if err != nil {
			return err
		}
		handler(m)
	}
}
