package events

import (
	"solid-example-go/internal/domain/book"
	"time"
)

type BookUpdatedDomainEvent struct {
	UpdatedDate time.Time
	FormerBook  book.Book
	NewBook     book.Book
}

func NewBookUpdatedDomainEvent(date time.Time, former, new book.Book) BookUpdatedDomainEvent {
	return BookUpdatedDomainEvent{
		UpdatedDate: date,
		FormerBook:  former,
		NewBook:     new,
	}
}

func (b BookUpdatedDomainEvent) EventName() string {
	return "book_updated"
}
