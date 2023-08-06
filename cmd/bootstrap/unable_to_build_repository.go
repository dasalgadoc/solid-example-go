package bootstrap

import "fmt"

type UnableToBuildRepository struct {
	Message string
}

func BuildRepositoryError(repoName string) *UnableToBuildRepository {
	return &UnableToBuildRepository{
		Message: fmt.Sprintf("unable to build %s repository", repoName),
	}
}

func (e *UnableToBuildRepository) Error() string {
	return e.Message
}
