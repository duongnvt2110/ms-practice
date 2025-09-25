package main

import (
	"booking-service/pkg/config"
	"fmt"
	"net"
	"strconv"

	"github.com/davecgh/go-spew/spew"
	"github.com/segmentio/kafka-go"
)

func main() {
	createTopic()
}

func createTopic() {
	// to create topics when auto.create.topics.enable='false'
	topic := "booking.events"

	cfg := config.NewConfig()
	conn, err := kafka.Dial("tcp", cfg.Kafka.Brokers[0])
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	spew.Dump(conn.Controller())

	controller, err := conn.Controller()
	if err != nil {
		panic(err.Error())
	}
	var controllerConn *kafka.Conn
	fmt.Println(controller)
	controllerConn, err = kafka.Dial("tcp", net.JoinHostPort(controller.Host, strconv.Itoa(controller.Port)))
	if err != nil {
		panic(err.Error())
	}
	defer controllerConn.Close()

	topicConfigs := []kafka.TopicConfig{
		{
			Topic:             topic,
			NumPartitions:     1,
			ReplicationFactor: 1,
		},
	}

	err = controllerConn.CreateTopics(topicConfigs...)
	if err != nil {
		panic(err.Error())
	}
}
