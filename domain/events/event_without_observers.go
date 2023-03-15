package events

import "fmt"

type EventWithoutObservers struct {
	Message string
}

func EventNotFound(eventName string) *EventWithoutObservers {
	return &EventWithoutObservers{
		Message: fmt.Sprintf("event %s has no observers registred", eventName),
	}
}

func (e *EventWithoutObservers) Error() string {
	return e.Message
}
