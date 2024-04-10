package akafka

import "github.com/confluentinc/confluent-kafka-go/kafka"

type KafkaProducer struct {
	KafkaProducer *kafka.Producer
}

func NewKafkaProducer(server, groupId string) *KafkaProducer {
	kafkaProducer, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": server,
	})
	if err != nil {
		panic(err)
	}
	return &KafkaProducer{
		KafkaProducer: kafkaProducer,
	}
}

func (k *KafkaProducer) ProduceMessage(topic string, message []byte) error {
	deliveryChan := make(chan kafka.Event)
	err := k.KafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic, Partition: -1},
		Value:          message,
	}, deliveryChan)
	if err != nil {
		return err
	}
	e := <-deliveryChan
	m := e.(*kafka.Message)
	if m.TopicPartition.Error != nil {
		return m.TopicPartition.Error
	}
	return nil
}
