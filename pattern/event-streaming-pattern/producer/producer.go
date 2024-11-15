package producer

import (
	"encoding/json"
	"event-streaming-pattern/domain"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

type EventProducer struct {
	natsConn *nats.Conn
	subject  string
}

func NewEventProducer(natsURL, subject string) (*EventProducer, error) {
	conn, err := nats.Connect(natsURL)
	if err != nil {
		return nil, err
	}

	return &EventProducer{
		natsConn: conn,
		subject:  subject,
	}, nil
}

func (p *EventProducer) PublishEvent(event domain.Event) error {
	eventData, err := json.Marshal(event)
	if err != nil {
		return fmt.Errorf("Faild to marshal event: %w", err)
	}

	err = p.natsConn.Publish(p.subject, eventData)
	if err != nil {
		return fmt.Errorf("Failed to publish event: %w", err)
	}
	log.Printf("Publish event: %s", event)
	return nil
}
