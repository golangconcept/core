package domain

import "fmt"

type Event struct {
	EventType string
	UserID    string
	UserName  string
	Email     string
}

func (e Event) String() string {
	return fmt.Sprintf("Event  Type: %s UserID: %s, UserName: %s, Email: %s", e.EventType, e.UserID, e.UserName, e.Email)
}
