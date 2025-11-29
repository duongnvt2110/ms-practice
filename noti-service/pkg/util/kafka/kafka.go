package kafka

import (
	"ms-practice/pkg/config"
	"ms-practice/pkg/event"
	shared "ms-practice/pkg/kafka"
	"sync"
)

type NotificationMessaging struct {
	Consumers map[string]shared.KafkaClient
	cfg       *config.Kafka
}

var (
	once      sync.Once
	messaging *NotificationMessaging
)

var topics = []event.Topic{
	{Name: event.PaymentTopicName, GroupID: "noti-payment-consumer", EnableConsumer: true},
	{Name: event.TicketTopicName, GroupID: "noti-ticket-consumer", EnableConsumer: true},
}

func NewNotificationMessaging(cfg *config.Kafka) *NotificationMessaging {
	if cfg == nil {
		return nil
	}
	once.Do(func() {
		messaging = &NotificationMessaging{
			Consumers: make(map[string]shared.KafkaClient),
			cfg:       cfg,
		}
		initialize(messaging)
	})
	return messaging
}

func initialize(m *NotificationMessaging) {
	for _, topic := range topics {
		if topic.EnableConsumer {
			group := topic.GroupID
			if group == "" {
				group = topic.Name + "-noti"
			}
			m.Consumers[topic.Name] = shared.NewKafkaClient(m.cfg).SetReaderTopic(topic.Name, group)
		}
	}
}
