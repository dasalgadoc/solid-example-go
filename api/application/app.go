package application

import (
	"solid-example-go/infrastructure/book"
	"solid-example-go/infrastructure/ping"
)

type (
	Application struct {
		ping        ping.Ping
		bookPrinter book.BookPrinterController
	}
)

func BuildApplication() *Application {
	return &Application{}
}
