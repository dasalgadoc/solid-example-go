package bootstrap

import (
	bookUC "solid-example-go/internal/application/book"
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/infrastructure/book/entrypoints"
)

func BuildBookFinder(repository book.BookRepository,
) entrypoints.HttpBookFinder {
	bookGetterUC := buildBookFinderUC(repository)
	return entrypoints.NewHttpBookFinder(bookGetterUC)
}

func buildBookFinderUC(repository book.BookRepository) bookUC.BookGetter {
	return bookUC.NewBookGetter(repository)
}
