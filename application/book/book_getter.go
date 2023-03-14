package book

import "solid-example-go/domain/book"

type BookGetter struct {
	bookRepository book.BookRepository
}

func NewBookGetter(r book.BookRepository) BookGetter {
	return BookGetter{
		bookRepository: r,
	}
}

func (bg *BookGetter) Find(ISBN string) (book.Book, error) {
	return bg.bookRepository.Find(ISBN)
}
