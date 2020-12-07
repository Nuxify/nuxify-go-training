package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "rest-server/configs/hystrix"
	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// PostCommandRepositoryCircuitBreaker circuit breaker for post command repository
type PostCommandRepositoryCircuitBreaker struct {
	repository.PostCommandRepositoryInterface
}

var config = hystrix_config.Config{}

// DeletePostByID is the decorator for the the post repository delete by id method
func (repository *PostCommandRepositoryCircuitBreaker) DeletePostByID(postID int64) error {
	hystrix.ConfigureCommand("delete_post_by_id", config.Settings())
	errors := hystrix.Go("delete_post_by_id", func() error {
		err := repository.PostCommandRepositoryInterface.DeletePostByID(postID)
		if err != nil {
			return err
		}

		return nil
	}, nil)

	select {
	case err := <-errors:
		return err
	default:
		return nil
	}
}

// InsertPost decorator pattern to insert post
func (repository *PostCommandRepositoryCircuitBreaker) InsertPost(data repositoryTypes.CreatePost) (entity.Post, error) {
	output := make(chan entity.Post, 1)
	hystrix.ConfigureCommand("insert_post", config.Settings())
	errors := hystrix.Go("insert_post", func() error {
		post, err := repository.PostCommandRepositoryInterface.InsertPost(data)
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

// UpdatePostByID is the decorator for the post repository update post method
func (repository *PostCommandRepositoryCircuitBreaker) UpdatePostByID(data repositoryTypes.UpdatePost) (entity.Post, error) {
	output := make(chan entity.Post, 1)
	hystrix.ConfigureCommand("update_post", config.Settings())
	errors := hystrix.Go("update_post", func() error {
		post, err := repository.PostCommandRepositoryInterface.UpdatePostByID(data)
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
