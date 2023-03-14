package providers

import (
	"solid-example-go/api/configs"
	"solid-example-go/domain/book"
	"solid-example-go/infrastructure/book/repositories"
)

const repositoryName = "book"

var bookRepositories = map[string]func(config configs.Config) (book.BookRepository, error){
	configs.InMemoryRepoType: buildInMemoryBookRepository,
	configs.MySqlRepoType:    buildMySQLBookRepository,
}

func GetBookRepository(config configs.Config) (book.BookRepository, error) {
	bookRepository, exists := bookRepositories[config.Repository.Type]
	if !exists {
		return nil, BuildRepositoryError(repositoryName)
	}

	return bookRepository(config)
}

func buildInMemoryBookRepository(config configs.Config) (book.BookRepository, error) {
	return repositories.NewInMemoryBookRepository()
}

func buildMySQLBookRepository(config configs.Config) (book.BookRepository, error) {
	return repositories.NewMySQLBookRepository()
}
