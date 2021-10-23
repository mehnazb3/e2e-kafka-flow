package producer

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

func Init() {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"acks":              "all",
	})
	if err != nil {
		panic(err)
	}
	defer kafkaProducer.Close()
	topic := "first_topic"
	msg := kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: kafka.PartitionAny},
		Value:          []byte("Hello Mehnaz"),
		Key:            []byte("Id1"),
	}
	deliveryChan := make(chan kafka.Event)
	kafkaProducer.Produce(&msg, deliveryChan)

	ev := <-deliveryChan
	mes := ev.(*kafka.Message)

	fmt.Println("msg, ", mes)
}
