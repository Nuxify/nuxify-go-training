package application

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/infrastructure/service/types"
)

// PostCommandServiceInterface holds the implementable method for the Post command service
type PostCommandServiceInterface interface {
	CreatePost(ctx context.Context, data types.CreatePost) (entity.Post, error)
	DeletePostByID(postID int64) error
	UpdatePostByID(ctx context.Context, data types.UpdatePost) (entity.Post, error)
}
