package events

import "fmt"

type EventUnconvertible struct {
	Message string
}

func EventCastFail(eventSource, eventTarget string) *EventUnconvertible {
	return &EventUnconvertible{
		Message: fmt.Sprintf("event %s cant be converted to %s", eventSource, eventTarget),
	}
}

func (e *EventUnconvertible) Error() string {
	return e.Message
}
