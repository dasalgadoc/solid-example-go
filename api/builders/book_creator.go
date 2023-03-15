package builders

import (
	bookUC "solid-example-go/application/book"
	"solid-example-go/domain/book"
	"solid-example-go/domain/events"
	"solid-example-go/infrastructure/book/entrypoints"
)

func BuildBookCreator(repository book.BookRepository,
	bus events.EventBus,
) entrypoints.HttpBookCreator {
	bookCreatorUC := buildBookCreatorUC(repository, bus)
	return entrypoints.NewHttpBookCreator(bookCreatorUC)
}

func buildBookCreatorUC(repository book.BookRepository,
	bus events.EventBus,
) bookUC.BookCreator {
	return bookUC.NewBookCreator(repository, bus)
}
