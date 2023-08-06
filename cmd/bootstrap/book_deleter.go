package bootstrap

import (
	bookUC "solid-example-go/internal/application/book"
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/infrastructure/book/entrypoints"
)

func BuildBookDeleter(repository book.BookRepository,
) entrypoints.HttpBookDeleter {
	bookDeleterUC := buildBookDeleterUC(repository)
	return entrypoints.NewHttpBookDeleter(bookDeleterUC)
}

func buildBookDeleterUC(repository book.BookRepository) bookUC.BookDeleter {
	return bookUC.NewBookDeleter(repository)
}
