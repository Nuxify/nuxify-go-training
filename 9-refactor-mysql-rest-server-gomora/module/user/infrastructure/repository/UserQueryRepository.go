package repository

import (
	"errors"
	"fmt"

	"rest-server/infrastructures/database/mysql/types"
	apiError "rest-server/internal/errors"
	"rest-server/module/user/domain/entity"
	repositoryTypes "rest-server/module/user/infrastructure/repository/types"
)

// UserQueryRepository handles databas access logic
type UserQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectUsers select a users
func (repository *UserQueryRepository) SelectUsers() ([]entity.User, error) {
	var user entity.User
	var users []entity.User

	stmt := fmt.Sprintf("SELECT * FROM %s", user.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{}, &users)
	if err != nil {
		return users, errors.New(apiError.DatabaseError)
	} else if len(users) == 0 {
		return users, errors.New(apiError.MissingRecord)
	}

	return users, nil
}

// SelectUserByID select a user by user id
func (repository *UserQueryRepository) SelectUserByID(data repositoryTypes.GetUser) (entity.User, error) {
	var user entity.User
	var users []entity.User

	stmt := fmt.Sprintf("SELECT * FROM %s Where id=:id", user.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{"id": data.ID}, &users)
	if err != nil {
		return entity.User{}, errors.New(apiError.DatabaseError)
	} else if len(users) == 0 {
		return entity.User{}, errors.New(apiError.MissingRecord)
	}

	return users[0], nil
}
