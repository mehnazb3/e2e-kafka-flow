package main

import (
	"fmt"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

var timeDuration time.Duration = 3000

func main() {
	newConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers":        "localhost:9092",
		"group.id":                 "e2e-kafka-group",
		"auto.offset.reset":        "earliest",
		"enable.auto.offset.store": true,
		// "debug":             "broker,topic,queue,msg,protocol",
	})
	if err != nil {
		fmt.Println("Err", err)
		return
	}
	defer newConsumer.Close()
	newConsumer.SubscribeTopics([]string{"first_topic"}, nil)

	for {
		msg, err := newConsumer.ReadMessage(timeDuration)
		if err != nil {
			fmt.Println("Err", err)
			return
		}
		fmt.Println(msg.Value)
		fmt.Println(msg.String())
		fmt.Println(msg.TopicPartition)
	}

}
