package events

type Producer string
type Consumer string

const (
	BookingTopicName     Producer = "booking.events"
	BoookingConsumerName Consumer = "booking.consumer"
)

type Topic struct {
	ProducerName Producer
	ConsumerName Consumer
	GroupID      string
	Enable       bool
}

var BookingTopic = []Topic{
	{
		ProducerName: BookingTopicName,
		ConsumerName: BoookingConsumerName,
		GroupID:      "booking.consumer",
		Enable:       true,
	},
}
