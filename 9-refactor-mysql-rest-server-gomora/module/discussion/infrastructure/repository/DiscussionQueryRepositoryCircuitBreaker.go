package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
)

// PostQueryRepositoryCircuitBreaker is the circuit breaker for the post query repository
type PostQueryRepositoryCircuitBreaker struct {
	repository.DiscussionQueryRepositoryInterface
}

// CommentQueryRepositoryCircuitBreaker is the circuit breaker for the comment query repository
type CommentQueryRepositoryCircuitBreaker struct {
	repository.DiscussionQueryRepositoryInterface
}

// =====================================POST=====================================

// SelectPosts is a decorator for the select posts repository
func (repository *PostQueryRepositoryCircuitBreaker) SelectPosts() ([]entity.Post, error) {
	output := make(chan []entity.Post, 1)
	hystrix.ConfigureCommand("select_post", config.Settings())
	errors := hystrix.Go("select_post", func() error {
		posts, err := repository.DiscussionQueryRepositoryInterface.SelectPosts()
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

// =====================================COMMENT=====================================

// SelectComments is a decorator for the select comments repository
func (repository *CommentQueryRepositoryCircuitBreaker) SelectComments() ([]entity.Comment, error) {
	output := make(chan []entity.Comment, 1)
	hystrix.ConfigureCommand("select_comment", config.Settings())
	errors := hystrix.Go("select_comment", func() error {
		comments, err := repository.DiscussionQueryRepositoryInterface.SelectComments()
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
