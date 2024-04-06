package main

import (
	"github.com/kzar1n/events-consumer/internal/infra/akafka"
	"github.com/kzar1n/events-consumer/internal/infra/database"
)

func main() {

	database.InitializeDatabase()
	kafkaConsumer := akafka.NewKafkaConsumer()
	go kafkaConsumer.SubscribeTopics([]string{"events"})

	kafkaConsumer.ConsumeMessages()

}
