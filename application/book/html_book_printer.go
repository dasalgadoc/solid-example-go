package book

import (
	"fmt"
	"solid-example-go/domain/book"
)

type HtmlBookPrinter struct{}

func NewHtmlBookPrinter() HtmlBookPrinter {
	return HtmlBookPrinter{}
}

func (hbp *HtmlBookPrinter) FormatBook(book book.Book) string {
	return fmt.Sprintf("<h1><strong>%s</strong> - %s</h1></br><p>%s</p>",
		book.Author, book.Title, book.CurrentPage)
}

func (hbp *HtmlBookPrinter) ServeBook(title, author string) string {
	book := book.NewBook("", title, author, "")
	return hbp.FormatBook(book)
}
