package kafka

import (
	"ms-practice/pkg/config"
	"ms-practice/pkg/event"
	sharedKafaka "ms-practice/pkg/kafka"
	"sync"
)

var (
	cfgOnce          sync.Once
	bookingMessaging *BookingMessaging
)

var topics = []event.Topic{
	{
		Name:           event.BookingTopicName,
		EnableProducer: true,
	},
	{
		Name:           event.PaymentTopicName,
		GroupID:        "booking-payment-consumer",
		EnableConsumer: true,
	},
}

type BookingMessaging struct {
	Consumers map[string]sharedKafaka.KafkaClient
	Producers map[string]sharedKafaka.KafkaClient
	cfg       *config.Kafka
}

func NewBookingKafkaClient(cfg *config.Kafka) *BookingMessaging {
	bookingMessaging = &BookingMessaging{
		Consumers: make(map[string]sharedKafaka.KafkaClient),
		Producers: make(map[string]sharedKafaka.KafkaClient),
		cfg:       cfg,
	}
	cfgOnce.Do(func() {
		initalizeKafkaConnection(bookingMessaging)
	})
	return bookingMessaging
}

func initalizeKafkaConnection(k *BookingMessaging) {
	initializeProducer(k)
	initializeConsumer(k)
}

func initializeProducer(k *BookingMessaging) {
	for _, topic := range topics {
		if topic.EnableProducer {
			client := sharedKafaka.NewKafkaClient(k.cfg).SetWriterTopic(topic.Name)
			k.Producers[topic.Name] = client
		}
	}
}

func initializeConsumer(k *BookingMessaging) {
	for _, topic := range topics {
		if topic.EnableConsumer {
			group := topic.GroupID
			if group == "" {
				group = topic.Name + "-group"
			}
			client := sharedKafaka.NewKafkaClient(k.cfg).SetReaderTopic(topic.Name, group)
			k.Consumers[topic.Name] = client
		}
	}
}
