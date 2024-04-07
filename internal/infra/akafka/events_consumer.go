package akafka

import (
	"encoding/json"
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/kzar1n/events-consumer/internal/services"
)

const (
	KafkaServer  = "host.docker.internal:9092"
	KafkaGroupID = "group_id"
)

type KafkaConsumer struct {
	Consumer *kafka.Consumer
	msgChan  chan *kafka.Message
}

func NewKafkaConsumer() *KafkaConsumer {
	msgChan := make(chan *kafka.Message)
	kafkaConsumer, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": KafkaServer,
		"group.id":          KafkaGroupID,
		"auto.offset.reset": "earliest",
	})
	if err != nil {
		panic(err)
	}
	return &KafkaConsumer{
		Consumer: kafkaConsumer,
		msgChan:  msgChan,
	}
}

func (k *KafkaConsumer) SubscribeTopics(topics []string) {
	if err := k.Consumer.SubscribeTopics(topics, nil); err != nil {
		panic(err)
	} else {
		fmt.Println("Subscribed to topics: ", topics)
	}

	for {
		if msg, err := k.Consumer.ReadMessage(-1); err == nil {
			k.msgChan <- msg
		}
	}
}

func (k *KafkaConsumer) ConsumeMessages() {
	for msg := range k.msgChan {
		eventDto := services.EventDTO{}

		if err := json.Unmarshal(msg.Value, &eventDto); err != nil {
			fmt.Println("Invalid event: ", msg.Value)
		}

		if eventDto.IdPayment != "" {
			eventDto.EventProcessor()
		} else {
			fmt.Println("Invalid event: ", msg.Value)
		}
	}
}
