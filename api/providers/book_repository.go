package providers

import (
	"solid-example-go/domain/book"
	"solid-example-go/infrastructure/book/repositories"
)

func GetBookRepository() (book.BookRepository, error) {
	return buildInMemoryBookRepository()
}

func buildInMemoryBookRepository() (book.BookRepository, error) {
	return repositories.NewInMemoryBookRepository(), nil
}
