package book

import "solid-example-go/domain/book"

type BookCreator struct {
	bookRepository book.BookRepository
}

func NewBookCreator(r book.BookRepository) BookCreator {
	return BookCreator{
		bookRepository: r,
	}
}

func (bc *BookCreator) Create(ISBN, author, title string) error {
	book := book.NewBook(ISBN, title, author, "")
	return bc.bookRepository.Save(book)
}
