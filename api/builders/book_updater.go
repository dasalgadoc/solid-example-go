package builders

import (
	bookUC "solid-example-go/application/book"
	"solid-example-go/domain/book"
	"solid-example-go/domain/events"
	"solid-example-go/infrastructure/book/entrypoints"
)

func BuildBookUpdater(repository book.BookRepository,
	bus events.EventBus,
) entrypoints.HttpBookUpdater {
	bookUpdaterUc := buildBookUpdaterUC(repository, bus)
	return entrypoints.NewBookUpdater(bookUpdaterUc)
}

func buildBookUpdaterUC(repository book.BookRepository, bus events.EventBus) bookUC.BookUpdater {
	return bookUC.NewBookUpdater(repository, bus)
}
