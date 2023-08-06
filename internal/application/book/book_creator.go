package book

import (
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/domain/events"
)

type BookCreator struct {
	bookRepository book.BookRepository
	eventBus       events.EventBus
}

func NewBookCreator(r book.BookRepository, e events.EventBus) BookCreator {
	return BookCreator{
		bookRepository: r,
		eventBus:       e,
	}
}

func (bc *BookCreator) Create(ISBN, author, title string) error {
	book := book.NewBook(ISBN, title, author, "")
	err := bc.bookRepository.Save(book)
	if err != nil {
		return err
	}

	return bc.eventBus.Notify(events.NewBookCreatedDomainEvent(ISBN))
}
