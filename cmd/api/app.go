package api

import (
	"fmt"
	"solid-example-go/cmd/bootstrap"
	"solid-example-go/cmd/configs"
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/domain/events"
	"solid-example-go/internal/infrastructure/book/entrypoints"
	"solid-example-go/internal/infrastructure/ping"
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
		bookCreator: bootstrap.BuildBookCreator(repositories.bookRepository, eventBus.eventbus),
		bookGetter:  bootstrap.BuildBookFinder(repositories.bookRepository),
		bookUpdater: bootstrap.BuildBookUpdater(repositories.bookRepository, eventBus.eventbus),
		bookDeleter: bootstrap.BuildBookDeleter(repositories.bookRepository),
	}
}

func getConfiguration() configs.Config {
	appConfig, err := configs.LoadConfig("./cmd/configs/config.yaml")
	if err != nil {
		panic(fmt.Errorf("error getting configuration: %w", err))
	}
	return appConfig
}

func buildRepositories(config configs.Config) (*applicationRepositories, error) {
	books, err := bootstrap.GetBookRepository(config)
	if err != nil {
		return nil, fmt.Errorf("error getting provider BookRepository: %w", err)
	}
	return &applicationRepositories{
		bookRepository: books,
	}, nil
}

func buildEventBus(config configs.Config) (*applicationEventBus, error) {
	bus, err := bootstrap.GetEventBus(config)
	if err != nil {
		return nil, fmt.Errorf("error getting provider EventBus: %w", err)
	}
	return &applicationEventBus{
		eventbus: bus,
	}, nil
}
