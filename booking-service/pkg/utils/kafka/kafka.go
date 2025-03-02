package kafka_client

import (
	"booking-service/pkg/config"

	"github.com/segmentio/kafka-go"
)

type KafkaClient interface {
}

type kafkaClient struct {
	Writer *kafka.Writer
	Reader *kafka.Reader
}

func NewKafkaClient(cfg *config.Config) KafkaClient {
	kWriter := kafka.NewWriter(kafka.WriterConfig{
		Brokers: cfg.Kafka.Brokers,
	})
	kReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: cfg.Kafka.Brokers,
	})
	return kafkaClient{
		Writer: kWriter,
		Reader: kReader,
	}
}

func (k *kafkaClient) SetWriterTopic(topic string) KafkaClient {
	k.Writer.Topic = topic
	return k
}

func (k *kafkaClient) SetReaderTopic(topic string, groupId string) KafkaClient {
	kCfg := k.Reader.Config()
	kReader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: kCfg.Brokers,
		Topic:   topic,
		GroupID: groupId,
	})
	k.Reader = kReader
	return k
}
