package book

type BookRepository interface {
	Save(b Book) error
	Find(ISBN string) (Book, error)
	Update(b Book) error
	Delete(ISBN string) error
}
