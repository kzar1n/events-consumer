package main

import (
	"github.com/kzar1n/events-consumer/internal/infra/akafka"
	"github.com/kzar1n/events-consumer/internal/infra/database"
)

func main() {

	database.ConnectMysqlDB()
	database.ConnectPostgreDB()
	kafkaConsumer := akafka.NewKafkaConsumer("host.docker.internal:9092", "group_id")
	go kafkaConsumer.SubscribeTopics([]string{"events"})

	kafkaConsumer.ConsumeMessages()

	defer database.MysqlDB.Close()
	defer database.PostgreDB.Close()
}
