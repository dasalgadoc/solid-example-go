package repositories

import (
	"fmt"
	"solid-example-go/domain/book"
)

type MySQLBookRepository struct {
	// details like connection string or whatever
}

func NewMySQLBookRepository() (book.BookRepository, error) {
	return &MySQLBookRepository{}, nil
}

func (m *MySQLBookRepository) Save(b book.Book) error {
	return fmt.Errorf("Not implemented yet")
}

func (m *MySQLBookRepository) Find(ISBN string) (book.Book, error) {
	return book.Book{}, fmt.Errorf("Not implemented yet")
}

func (m *MySQLBookRepository) Update(b book.Book) error {
	return fmt.Errorf("Not implemented yet")
}

func (m *MySQLBookRepository) Delete(ISB string) error {
	return fmt.Errorf("Not implemented yet")
}
