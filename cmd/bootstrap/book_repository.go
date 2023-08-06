package bootstrap

import (
	"solid-example-go/cmd/configs"
	"solid-example-go/internal/domain/book"
	"solid-example-go/internal/infrastructure/book/repositories"
)

const repositoryName = "book"

var bookRepositories = map[string]func(config configs.Config) (book.BookRepository, error){
	configs.InMemoryType:  buildInMemoryBookRepository,
	configs.MySqlRepoType: buildMySQLBookRepository,
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
