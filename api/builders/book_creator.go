package builders

import (
	bookUC "solid-example-go/application/book"
	"solid-example-go/domain/book"
	"solid-example-go/infrastructure/book/entrypoints"
)

func BuildBookCreator(repository book.BookRepository,
) entrypoints.HttpBookCreator {
	bookCreatorUC := buildBookCreatorUC(repository)
	return entrypoints.NewHttpBookCreator(bookCreatorUC)
}

func buildBookCreatorUC(repository book.BookRepository) bookUC.BookCreator {
	return bookUC.NewBookCreator(repository)
}
