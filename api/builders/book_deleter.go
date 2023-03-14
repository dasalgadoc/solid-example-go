package builders

import (
	book2 "solid-example-go/application/book"
	"solid-example-go/domain/book"
	"solid-example-go/infrastructure/book/entrypoints"
)

func BuildBookDeleter(repository book.BookRepository,
) entrypoints.HttpBookDeleter {
	bookDeleterUC := buildBookDeleterUC(repository)
	return entrypoints.NewHttpBookDeleter(bookDeleterUC)
}

func buildBookDeleterUC(repository book.BookRepository) book2.BookDeleter {
	return book2.NewBookDeleter(repository)
}
