package application

import (
	"fmt"
	"solid-example-go/api/builders"
	"solid-example-go/api/configs"
	"solid-example-go/api/providers"
	"solid-example-go/domain/book"
	"solid-example-go/domain/events"
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

	applicationEventBus struct {
		eventbus events.EventBus
	}
)

func BuildApplication() *Application {
	appConfig := getConfiguration()

	repositories, err := buildRepositories(appConfig)
	if err != nil {
		panic(fmt.Errorf("error building repositories: %w", err))
	}

	eventBus, err := buildEventBus(appConfig)
	if err != nil {
		panic(fmt.Errorf("error building event bus: %w", err))
	}

	return &Application{
		bookCreator: builders.BuildBookCreator(repositories.bookRepository, eventBus.eventbus),
		bookGetter:  builders.BuildBookFinder(repositories.bookRepository),
		bookUpdater: builders.BuildBookUpdater(repositories.bookRepository),
		bookDeleter: builders.BuildBookDeleter(repositories.bookRepository),
	}
}

func getConfiguration() configs.Config {
	appConfig, err := configs.LoadConfig("./api/configs/config.yaml")
	if err != nil {
		panic(fmt.Errorf("error getting configuration: %w", err))
	}
	return appConfig
}

func buildRepositories(config configs.Config) (*applicationRepositories, error) {
	books, err := providers.GetBookRepository(config)
	if err != nil {
		return nil, fmt.Errorf("error getting provider BookRepository: %w", err)
	}
	return &applicationRepositories{
		bookRepository: books,
	}, nil
}

func buildEventBus(config configs.Config) (*applicationEventBus, error) {
	bus, err := providers.GetEventBus(config)
	if err != nil {
		return nil, fmt.Errorf("error getting provider EventBus: %w", err)
	}
	return &applicationEventBus{
		eventbus: bus,
	}, nil
}
