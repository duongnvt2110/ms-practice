package kafka

import (
	"ms-practice/booking-service/pkg/events"
	"sync"

	sharedKafaka "ms-practice/pkg/kafka"
)

var (
	cfgOnce          sync.Once
	bookingMessaging *BookingMessaging
)

type BookingMessaging struct {
	Consumers map[events.Consumer]sharedKafaka.KafkaClient
	Producers map[events.Producer]sharedKafaka.KafkaClient
}

func NewBookingKafkaClient(kafkaClient sharedKafaka.KafkaClient) *BookingMessaging {
	bookingMessaging := &BookingMessaging{
		Consumers: make(map[events.Consumer]sharedKafaka.KafkaClient),
		Producers: make(map[events.Producer]sharedKafaka.KafkaClient),
	}
	cfgOnce.Do(func() {
		initalizeKafkaConnection(kafkaClient, bookingMessaging)
	})
	return bookingMessaging
}

func initalizeKafkaConnection(kafkaClient sharedKafaka.KafkaClient, k *BookingMessaging) {
	initializeProducer(kafkaClient, k)
	initializeConsumer(kafkaClient, k)
}

func initializeProducer(kafkaClient sharedKafaka.KafkaClient, k *BookingMessaging) {
	for _, topic := range events.BookingTopic {
		if topic.Enable {
			topicName := events.Producer(topic.ProducerName)
			k.Producers[topicName] = kafkaClient.SetWriterTopic(string(topicName))
		}
	}
}

func initializeConsumer(kafkaClient sharedKafaka.KafkaClient, k *BookingMessaging) {
	for _, topic := range events.BookingTopic {
		if topic.Enable {
			topicName := events.Producer(topic.ProducerName)
			consumerName := events.Consumer(topic.ConsumerName)
			k.Consumers[consumerName] = kafkaClient.SetReaderTopic(string(topicName), string(topic.GroupID))
		}
	}
}
