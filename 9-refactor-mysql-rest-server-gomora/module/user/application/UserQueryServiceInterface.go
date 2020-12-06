package application

import (
	"context"

	"rest-server/module/user/domain/entity"
	"rest-server/module/user/infrastructure/service/types"
)

// UserQueryServiceInterface holds the implementable method for the user query service
type UserQueryServiceInterface interface {
	GetUsers(ctx context.Context, data types.GetUser) ([]entity.User, error)
	GetUserByID(ctx context.Context, data types.GetUser) ([]entity.User, error)
}
