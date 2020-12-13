package application

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/infrastructure/service/types"
)

// DiscussionCommandServiceInterface holds the implementable method for the discussion command service
type DiscussionCommandServiceInterface interface {
	CreatePost(ctx context.Context, data types.CreatePost) (entity.Post, error)
	DeletePostByID(postID int64) error
	UpdatePostByID(ctx context.Context, data types.UpdatePost) (entity.Post, error)
	CreateComment(ctx context.Context, data types.CreateComment) (entity.Comment, error)
	DeleteCommentByID(commentID int64) error
	UpdateCommentByID(ctx context.Context, data types.UpdateComment) (entity.Comment, error)
}
