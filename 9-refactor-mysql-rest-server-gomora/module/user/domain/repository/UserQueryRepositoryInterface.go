package repository

import (
	"rest-server/module/user/domain/entity"
	repositoryTypes "rest-server/module/user/infrastructure/repository/types"
)

// UserQueryRepositoryInterface holds the methods for the user query repository
type UserQueryRepositoryInterface interface {
	SelectUsers() ([]entity.User, error)
	SelectUserByID(data repositoryTypes.GetUser) (entity.User, error)
}
