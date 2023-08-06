package book

import (
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/domain/events"
	"time"
)

type BookUpdater struct {
	bookRepository book.BookRepository
	eventBus       events.EventBus
}

func NewBookUpdater(r book.BookRepository, e events.EventBus) BookUpdater {
	return BookUpdater{
		bookRepository: r,
		eventBus:       e,
	}
}

func (bu *BookUpdater) Update(ISBN, author, title string) error {
	formerBook, err := bu.bookRepository.Find(ISBN)
	if err != nil {
		return book.DoesNotExist()
	}
	newBook := formerBook.ReplaceValues(author, title)
	err = bu.bookRepository.Update(newBook)
	if err != nil {
		return err
	}

	return bu.eventBus.Notify(events.NewBookUpdatedDomainEvent(time.Now(), formerBook, newBook))
}
