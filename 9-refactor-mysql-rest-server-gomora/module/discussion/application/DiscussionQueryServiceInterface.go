package application

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/infrastructure/service/types"
)

// DiscussionQueryServiceInterface holds the implementable method for the discussion query service
type DiscussionQueryServiceInterface interface {
	GetPosts(ctx context.Context) ([]entity.Post, error)
	GetPostByID(ctx context.Context, data types.GetPost) (entity.Post, error)
	GetComments(ctx context.Context) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, data types.GetComment) (entity.Comment, error)
}
