package book

import (
	"fmt"
	"solid-example-go/domain/book"
	"strings"
)

type ConsoleBookPrinter struct{}

func NewConsoleBookPrinter() ConsoleBookPrinter {
	return ConsoleBookPrinter{}
}

func (cbp *ConsoleBookPrinter) FormatBook(book book.Book) string {
	return fmt.Sprintf("%s - %s\n%s",
		strings.ToUpper(book.Author), book.Title, book.CurrentPage)
}

func (cbp *ConsoleBookPrinter) PrintBook(book book.Book) {
	fmt.Println(cbp.FormatBook(book))
}
