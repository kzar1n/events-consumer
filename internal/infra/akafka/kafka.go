package akafka

import (
	"fmt"

	"github.com/confluentinc/confluent-kafka-go/kafka"
)

const (
	// KafkaServer is the Kafka server address
	KafkaServer = "host.docker.internal:9092"
	// KafkaGroupID is the Kafka group ID
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
		fmt.Printf("Subscribed to topics: %v\n", topics)
	}

	for {
		msg, err := k.Consumer.ReadMessage(-1)
		if err == nil {
			k.msgChan <- msg
		}
	}
}

func (k *KafkaConsumer) ConsumeMessages() {
	for msg := range k.msgChan {
		fmt.Printf("Message on %s: %s\n", msg.TopicPartition, string(msg.Value))
		// Process message
	}
}

func SendMessage() {
	// Send message to Kafka
}
