package book

import (
	"github.com/stretchr/testify/assert"
	"solid-example-go/domain/book"
	"testing"
)

type consoleBookPrinterTestScenario struct {
	test               *testing.T
	consoleBookPrinter ConsoleBookPrinter
	book               book.Book
	result             string
}

func TestAnEmptyBook(t *testing.T) {
	s := startConsoleBookPrinterTestScenario(t)
	s.whenFormatBookIsRequired()
	s.thenExpectedFormatWasReturn(" - \n")
}

func TestABookWithAuthorOnly(t *testing.T) {
	s := startConsoleBookPrinterTestScenario(t)
	book := book.NewBook("", "", "John Doe", "")
	s.givenABook(book)
	s.whenFormatBookIsRequired()
	s.thenExpectedFormatWasReturn("JOHN DOE - \n")
}

func TestABookWithTitleOnly(t *testing.T) {
	s := startConsoleBookPrinterTestScenario(t)
	book := book.NewBook("", "1001 nights", "", "")
	s.givenABook(book)
	s.whenFormatBookIsRequired()
	s.thenExpectedFormatWasReturn(" - 1001 nights\n")
}

func TestABookWithCurrentPageOnly(t *testing.T) {
	s := startConsoleBookPrinterTestScenario(t)
	book := book.NewBook("", "", "", "page in blank")
	s.givenABook(book)
	s.whenFormatBookIsRequired()
	s.thenExpectedFormatWasReturn(" - \npage in blank")
}

func TestABookWithAuthorAndTitle(t *testing.T) {
	s := startConsoleBookPrinterTestScenario(t)
	book := book.NewBook("", "1001 nights", "john doe", "")
	s.givenABook(book)
	s.whenFormatBookIsRequired()
	s.thenExpectedFormatWasReturn("JOHN DOE - 1001 nights\n")
}

func TestAFullBook(t *testing.T) {
	s := startConsoleBookPrinterTestScenario(t)
	book := book.NewBook("", "1001 nights", "john doe", "page in blank")
	s.givenABook(book)
	s.whenFormatBookIsRequired()
	s.thenExpectedFormatWasReturn("JOHN DOE - 1001 nights\npage in blank")
}

/*-- steps --*/
func startConsoleBookPrinterTestScenario(t *testing.T) *consoleBookPrinterTestScenario {
	t.Parallel()
	return &consoleBookPrinterTestScenario{
		test:               t,
		consoleBookPrinter: NewConsoleBookPrinter(),
	}
}

func (c *consoleBookPrinterTestScenario) givenABook(book book.Book) {
	c.book = book
}

func (c *consoleBookPrinterTestScenario) whenFormatBookIsRequired() {
	c.result = c.consoleBookPrinter.FormatBook(c.book)
}

func (c *consoleBookPrinterTestScenario) thenExpectedFormatWasReturn(expected string) {
	assert.Equal(c.test, expected, c.result)
}
