package kafka

import (
	"ms-practice/payment-service/pkg/event"
	sharedKafaka "ms-practice/pkg/kafka"
	"sync"
)

var (
	cfgOnce          sync.Once
	bookingMessaging *BookingMessaging
)

type BookingMessaging struct {
	Consumers map[event.Consumer]sharedKafaka.KafkaClient
	Producers map[event.Producer]sharedKafaka.KafkaClient
}

func NewBookingKafkaClient(kafkaClient sharedKafaka.KafkaClient) *BookingMessaging {
	bookingMessaging = &BookingMessaging{
		Consumers: make(map[event.Consumer]sharedKafaka.KafkaClient),
		Producers: make(map[event.Producer]sharedKafaka.KafkaClient),
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
	for _, topic := range event.BookingTopic {
		if topic.Enable {
			topicName := event.Producer(topic.ProducerName)
			k.Producers[topicName] = kafkaClient.SetWriterTopic(string(topicName))
		}
	}
}

func initializeConsumer(kafkaClient sharedKafaka.KafkaClient, k *BookingMessaging) {
	for _, topic := range event.BookingTopic {
		if topic.Enable {
			topicName := event.Producer(topic.ProducerName)
			consumerName := event.Consumer(topic.ConsumerName)
			k.Consumers[consumerName] = kafkaClient.SetReaderTopic(string(topicName), string(topic.GroupID))
		}
	}
}
