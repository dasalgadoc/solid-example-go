package book

type Book struct {
	ISBN        string
	Title       string
	Author      string
	CurrentPage string
}

func NewBook(ISBN, title, author, currentPage string) Book {
	return Book{
		ISBN:        ISBN,
		Title:       title,
		Author:      author,
		CurrentPage: currentPage,
	}
}

func (b *Book) ReplaceValues(author, title string) Book {
	if author == "" {
		author = b.Author
	}
	if title == "" {
		title = b.Title
	}
	return Book{
		ISBN:        b.ISBN,
		Author:      author,
		Title:       title,
		CurrentPage: "",
	}
}
