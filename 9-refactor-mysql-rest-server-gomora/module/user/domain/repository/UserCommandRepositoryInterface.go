package repository

import (
	"rest-server/module/user/domain/entity"
	repositoryTypes "rest-server/module/user/infrastructure/repository/types"
)

// UserCommandRepositoryInterface holds the implementable methods for the academic year command repository
type UserCommandRepositoryInterface interface {
	DeleteUserByID(UserID int64) error
	InsertUser(data repositoryTypes.CreateUser) (entity.User, error)
	UpdateUserByID(data repositoryTypes.UpdateUser) (entity.User, error)
}
