package book

type BookFormatter interface {
	FormatBook(book Book) string
}
