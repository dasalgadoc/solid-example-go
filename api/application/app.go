package application

import (
	"fmt"
	"solid-example-go/api/builders"
	"solid-example-go/api/providers"
	"solid-example-go/domain/book"
	"solid-example-go/infrastructure/book/entrypoints"
	"solid-example-go/infrastructure/ping"
)

type (
	Application struct {
		ping        ping.Ping
		bookPrinter entrypoints.BookPrinterController
		bookCreator entrypoints.HttpBookCreator
		bookGetter  entrypoints.HttpBookFinder
		bookUpdater entrypoints.HttpBookUpdater
		bookDeleter entrypoints.HttpBookDeleter
	}

	applicationRepositories struct {
		bookRepository book.BookRepository
	}
)

func BuildApplication() *Application {
	repositories, err := buildRepositories()
	if err != nil {
		panic(fmt.Errorf("error building repositories: %w", err))
	}

	return &Application{
		bookCreator: builders.BuildBookCreator(repositories.bookRepository),
		bookGetter:  builders.BuildBookFinder(repositories.bookRepository),
		bookUpdater: builders.BuildBookUpdater(repositories.bookRepository),
		bookDeleter: builders.BuildBookDeleter(repositories.bookRepository),
	}
}

func buildRepositories() (*applicationRepositories, error) {
	books, err := providers.GetBookRepository()
	if err != nil {
		return nil, fmt.Errorf("error getting provider BookRepository: %w", err)
	}
	return &applicationRepositories{
		bookRepository: books,
	}, nil
}
