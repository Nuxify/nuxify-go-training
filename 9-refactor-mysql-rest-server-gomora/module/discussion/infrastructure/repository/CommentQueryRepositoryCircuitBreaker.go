package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// CommentQueryRepositoryCircuitBreaker is the circuit breaker for the comment query repository
type CommentQueryRepositoryCircuitBreaker struct {
	repository.CommentQueryRepositoryInterface
}

// SelectComments is a decorator for the select comments repository
func (repository *CommentQueryRepositoryCircuitBreaker) SelectComments(data repositoryTypes.GetComment) ([]entity.Comment, error) {
	output := make(chan []entity.Comment, 1)
	hystrix.ConfigureCommand("select_comment", config.Settings())
	errors := hystrix.Go("select_comment", func() error {
		comments, err := repository.CommentQueryRepositoryInterface.SelectComments(data)
		if err != nil {
			return err
		}

		output <- comments
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.Comment{}, err
	}
}
