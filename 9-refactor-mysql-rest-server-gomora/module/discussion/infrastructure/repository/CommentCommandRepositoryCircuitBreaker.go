package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	hystrix_config "rest-server/configs/hystrix"
	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// CommentCommandRepositoryCircuitBreaker circuit breaker for Comment command repository
type CommentCommandRepositoryCircuitBreaker struct {
	repository.CommentCommandRepositoryInterface
}

var conf = hystrix_config.Config{}

// DeleteCommentByID is the decorator for the the Comment repository delete by id method
func (repository *CommentCommandRepositoryCircuitBreaker) DeleteCommentByID(commentID int64) error {
	hystrix.ConfigureCommand("delete_comment_by_id", conf.Settings())
	errors := hystrix.Go("delete_comment_by_id", func() error {
		err := repository.CommentCommandRepositoryInterface.DeleteCommentByID(commentID)
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

// InsertComment decorator pattern to insert Comment
func (repository *CommentCommandRepositoryCircuitBreaker) InsertComment(data repositoryTypes.CreateComment) (entity.Comment, error) {
	output := make(chan entity.Comment, 1)
	hystrix.ConfigureCommand("insert_comment", conf.Settings())
	errors := hystrix.Go("insert_comment", func() error {
		comment, err := repository.CommentCommandRepositoryInterface.InsertComment(data)
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

// UpdateCommentByID is the decorator for the Comment repository update Comment method
func (repository *CommentCommandRepositoryCircuitBreaker) UpdateCommentByID(data repositoryTypes.UpdateComment) (entity.Comment, error) {
	output := make(chan entity.Comment, 1)
	hystrix.ConfigureCommand("update_comment", conf.Settings())
	errors := hystrix.Go("update_comment", func() error {
		comment, err := repository.CommentCommandRepositoryInterface.UpdateCommentByID(data)
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
