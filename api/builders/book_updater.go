package builders

import (
	bookUC "solid-example-go/application/book"
	"solid-example-go/domain/book"
	"solid-example-go/infrastructure/book/entrypoints"
)

func BuildBookUpdater(repository book.BookRepository,
) entrypoints.HttpBookUpdater {
	bookUpdaterUc := buildBookUpdaterUC(repository)
	return entrypoints.NewBookUpdater(bookUpdaterUc)
}

func buildBookUpdaterUC(repository book.BookRepository) bookUC.BookUpdater {
	return bookUC.NewBookUpdater(repository)
}
