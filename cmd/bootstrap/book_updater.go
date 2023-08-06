package bootstrap

import (
	bookUC "solid-example-go/internal/application/book"
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/domain/events"
	"solid-example-go/internal/infrastructure/book/entrypoints"
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
