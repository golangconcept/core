package consumer

import (
	"encoding/json"
	"event-streaming-pattern/domain"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type EventConsumer struct {
	natsConn *nats.Conn
	subject  string
}

func NewEventConsumer(natsURL, subject string) (*EventConsumer, error) {
	conn, err := nats.Connect(natsURL)

	if err != nil {
		return nil, err
	}

	return &EventConsumer{natsConn: conn, subject: subject}, err
}

func (c *EventConsumer) ListenForEvents() {
	_, err := c.natsConn.Subscribe(c.subject, func(msg *nats.Msg) {
		var event domain.Event
		err := json.Unmarshal(msg.Data, &event)

		if err != nil {
			log.Printf("Failed to unmarshal event: %v", err)
			return
		}

		fmt.Printf("Received event: %s\n", event)
	})
	if err != nil {
		log.Fatalf("Error subscribing to event: %v", err)
	}

	select {}
}
