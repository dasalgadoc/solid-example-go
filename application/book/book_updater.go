package book

import "solid-example-go/domain/book"

type BookUpdater struct {
	bookRepository book.BookRepository
}

func NewBookUpdater(r book.BookRepository) BookUpdater {
	return BookUpdater{
		bookRepository: r,
	}
}

func (bu *BookUpdater) Update(ISBN, author, title string) error {
	formerBook, _ := bu.bookRepository.Find(ISBN)
	newBook := formerBook.ReplaceValues(author, title)
	return bu.bookRepository.Update(newBook)
}
