package book

import (
	"solid-example-go/internal/domain/book"
)

type BookDeleter struct {
	bookRepository book.BookRepository
}

func NewBookDeleter(r book.BookRepository) BookDeleter {
	return BookDeleter{
		bookRepository: r,
	}
}

func (bd *BookDeleter) Delete(ISBN string) error {
	return bd.bookRepository.Delete(ISBN)
}
