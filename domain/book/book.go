package book

type Book struct {
	Title       string
	Author      string
	CurrentPage string
}

func NewBook(title, author, currentPage string) Book {
	return Book{
		Title:       title,
		Author:      author,
		CurrentPage: currentPage,
	}
}
