package events

import (
	"fmt"
)

type NewBookLogger struct{}

func (n *NewBookLogger) Update(event DomainEvent) error {
	bookCreatedDomainEvent, ok := event.(BookCreatedDomainEvent)
	if !ok {
		return EventCastFail(event.EventName(), "book_created")
	}
	log := fmt.Sprintf("---> NEW BOOK CREATED: %s", bookCreatedDomainEvent.ISBN)
	fmt.Println(log)
	return nil
}
