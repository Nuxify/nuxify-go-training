package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// PostQueryRepositoryCircuitBreaker is the circuit breaker for the post query repository
type PostQueryRepositoryCircuitBreaker struct {
	repository.PostQueryRepositoryInterface
}

// SelectPosts is a decorator for the select posts repository
func (repository *PostQueryRepositoryCircuitBreaker) SelectPosts(data repositoryTypes.GetPost) ([]entity.Post, error) {
	output := make(chan []entity.Post, 1)
	hystrix.ConfigureCommand("select_post", config.Settings())
	errors := hystrix.Go("select_post", func() error {
		posts, err := repository.PostQueryRepositoryInterface.SelectPosts(data)
		if err != nil {
			return err
		}

		output <- posts
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.Post{}, err
	}
}
