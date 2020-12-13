package repository

import (
	"github.com/afex/hystrix-go/hystrix"

	"rest-server/module/user/domain/entity"
	"rest-server/module/user/domain/repository"
	repositoryTypes "rest-server/module/user/infrastructure/repository/types"
)

// UserQueryRepositoryCircuitBreaker is the circuit breaker for the user query repository
type UserQueryRepositoryCircuitBreaker struct {
	repository.UserQueryRepositoryInterface
}

// SelectUsers is a decorator for the select users repository
func (repository *UserQueryRepositoryCircuitBreaker) SelectUsers() ([]entity.User, error) {
	output := make(chan []entity.User, 1)
	hystrix.ConfigureCommand("select_user", config.Settings())
	errors := hystrix.Go("select_user", func() error {
		users, err := repository.UserQueryRepositoryInterface.SelectUsers()
		if err != nil {
			return err
		}

		output <- users
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return []entity.User{}, err
	}
}

// SelectUserByID is a decorator for the select user by id repository
func (repository *UserQueryRepositoryCircuitBreaker) SelectUserByID(data repositoryTypes.GetUser) (entity.User, error) {
	output := make(chan entity.User, 1)
	hystrix.ConfigureCommand("select_user_by_id", config.Settings())
	errors := hystrix.Go("select_user_by_id", func() error {
		user, err := repository.UserQueryRepositoryInterface.SelectUserByID(data)
		if err != nil {
			return err
		}

		output <- user
		return nil
	}, nil)

	select {
	case out := <-output:
		return out, nil
	case err := <-errors:
		return entity.User{}, err
	}
}
