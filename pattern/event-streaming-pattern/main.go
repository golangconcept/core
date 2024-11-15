package main

import (
	"event-streaming-pattern/consumer"
	"event-streaming-pattern/domain"
	"event-streaming-pattern/producer"
	"fmt"
	"log"
	"time"
)

func main() {
	natsURL := "nats://localhost:4222"
	subject := "user.registration"

	eventConsumer, err := consumer.NewEventConsumer(natsURL, subject)

	if err != nil {
		log.Fatalf("Error creating event consumer: %v", err)
	}

	go eventConsumer.ListenForEvents()

	eventProducer, err := producer.NewEventProducer(natsURL, subject)

	if err != nil {
		log.Fatalf("Error creating event producer: %v", err)
	}

	for i := 1; i <= 5; i++ {
		event := domain.Event{
			EventType: "UserRegistration",
			UserID:    fmt.Sprintf("user%d", i),
			UserName:  fmt.Sprintf("User %d", i),
			Email:     fmt.Sprintf("user%demample.com", i),
		}

		err := eventProducer.PublishEvent(event)

		if err != nil {
			log.Printf("Error publishing event: %v", err)
		}
		time.Sleep(1 * time.Second)
	}
	time.Sleep(10 * time.Second)
}
