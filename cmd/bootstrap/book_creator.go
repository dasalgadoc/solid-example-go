package bootstrap

import (
	bookUC "solid-example-go/internal/application/book"
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/domain/events"
	"solid-example-go/internal/infrastructure/book/entrypoints"
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
