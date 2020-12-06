package application

import (
	"context"

	"rest-server/module/user/domain/entity"
	"rest-server/module/user/infrastructure/service/types"
)

// UserCommandServiceInterface holds the implementable method for the user command service
type UserCommandServiceInterface interface {
	CreateUser(ctx context.Context, data types.CreateUser) (entity.User, error)
	DeleteUserByID(userID int64) error
	UpdateUserByID(ctx context.Context, data types.UpdateUser) (entity.User, error)
}
