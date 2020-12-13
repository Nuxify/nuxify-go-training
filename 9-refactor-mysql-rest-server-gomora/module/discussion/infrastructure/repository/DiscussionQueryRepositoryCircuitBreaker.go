package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// DiscussionQueryRepositoryCircuitBreaker is the circuit breaker for the discussion query repository
type DiscussionQueryRepositoryCircuitBreaker struct {
	repository.DiscussionQueryRepositoryInterface
}

// =====================================POST=====================================

// SelectPosts is a decorator for the select posts repository
func (repository *DiscussionQueryRepositoryCircuitBreaker) SelectPosts() ([]entity.Post, error) {
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

// SelectPostByID is a decorator for the select post by id repository
func (repository *DiscussionQueryRepositoryCircuitBreaker) SelectPostByID(data repositoryTypes.GetPost) (entity.Post, error) {
	output := make(chan entity.Post, 1)
	hystrix.ConfigureCommand("select_post_by_id", config.Settings())
	errors := hystrix.Go("select_post_by_id", func() error {
		post, err := repository.DiscussionQueryRepositoryInterface.SelectPostByID(data)
		if err != nil {
			return err
		}

		output <- post
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Post{}, err
	}
}

// =====================================COMMENT=====================================

// SelectComments is a decorator for the select comments repository
func (repository *DiscussionQueryRepositoryCircuitBreaker) SelectComments() ([]entity.Comment, error) {
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

// SelectCommentByID is a decorator for the select comment by id repository
func (repository *DiscussionQueryRepositoryCircuitBreaker) SelectCommentByID(data repositoryTypes.GetComment) (entity.Comment, error) {
	output := make(chan entity.Comment, 1)
	hystrix.ConfigureCommand("select_comment_by_id", config.Settings())
	errors := hystrix.Go("select_comment_by_id", func() error {
		comment, err := repository.DiscussionQueryRepositoryInterface.SelectCommentByID(data)
		if err != nil {
			return err
		}

		output <- comment
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.Comment{}, err
	}
}
