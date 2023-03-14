package builders

import (
	bookUC "solid-example-go/application/book"
	"solid-example-go/domain/book"
	"solid-example-go/infrastructure/book/entrypoints"
)

func BuildBookFinder(repository book.BookRepository,
) entrypoints.HttpBookFinder {
	bookGetterUC := buildBookFinderUC(repository)
	return entrypoints.NewHttpBookFinder(bookGetterUC)
}

func buildBookFinderUC(repository book.BookRepository) bookUC.BookGetter {
	return bookUC.NewBookGetter(repository)
}
