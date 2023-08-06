package repositories

import (
	"solid-example-go/internal/domain/book"
)

type InMemoryBookRepository struct {
	Books map[string]book.Book
}

func NewInMemoryBookRepository() (book.BookRepository, error) {
	return &InMemoryBookRepository{
		Books: make(map[string]book.Book),
	}, nil
}

func (m *InMemoryBookRepository) Save(b book.Book) error {
	if m.bookExists(b.ISBN) {
		return book.AlreadyExists()
	}
	m.Books[b.ISBN] = b
	return nil
}

func (m *InMemoryBookRepository) Find(ISBN string) (book.Book, error) {
	return m.Books[ISBN], nil
}

func (m *InMemoryBookRepository) Update(b book.Book) error {
	if !m.bookExists(b.ISBN) {
		return book.DoesNotExist()
	}
	m.Books[b.ISBN] = b
	return nil
}

func (m *InMemoryBookRepository) Delete(ISBN string) error {
	if !m.bookExists(ISBN) {
		return book.DoesNotExist()
	}
	delete(m.Books, ISBN)
	return nil
}

func (m *InMemoryBookRepository) bookExists(ISBN string) bool {
	_, ok := m.Books[ISBN]
	return ok
}
