package main

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/kzar1n/events-consumer/internal/entity"
	"github.com/kzar1n/events-consumer/internal/infra/akafka"
	"github.com/kzar1n/events-consumer/internal/infra/database"
	"github.com/kzar1n/events-consumer/internal/services"
)

func main() {
	database.ConnectPostgreDB()
	kafkaProducer := akafka.NewKafkaProducer("host.docker.internal:9092", "group_id")
	for {

		idPayment := CreatePayment()

		event := CreateEvent(idPayment)

		eventJSON, err := json.Marshal(event)
		if err != nil {
			log.Fatalf("Failed to serialize event: %v", err)
		}

		kafkaProducer.ProduceMessage("events", eventJSON)
		fmt.Println("Event sent to Kafka: ", event)

		time.Sleep(1 * time.Second)
	}
	defer database.PostgreDB.Close()
}

func CreatePayment() string {
	idPayment := uuid.New().String()
	idAccount := uuid.New().String()
	Legacy := entity.NewLegacyPayment(idPayment, idAccount, "CREDIT")

	paymentService := services.NewPaymentService()
	if err := paymentService.SaveLegacyPayment(Legacy); err != nil {
		panic(err)
	}
	return idPayment
}

func CreateEvent(idPayment string) *services.EventDTO {
	id := uuid.New().String()
	eventType := "PAYMENT"
	eventDate := time.Now()
	event := entity.NewEvent(id, idPayment, eventType, eventDate)
	eventDTO := services.NewEventDTO(event)
	return eventDTO
}
