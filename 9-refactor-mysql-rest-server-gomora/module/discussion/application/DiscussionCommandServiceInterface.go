package application

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/infrastructure/service/types"
)

// PostCommandServiceInterface holds the implementable method for the post command service
type PostCommandServiceInterface interface {
	CreatePost(ctx context.Context, data types.CreatePost) (entity.Post, error)
	DeletePostByID(postID int64) error
	UpdatePostByID(ctx context.Context, data types.UpdatePost) (entity.Post, error)
}

// CommentCommandServiceInterface holds the implementable method for the comment command service
type CommentCommandServiceInterface interface {
	CreateComment(ctx context.Context, data types.CreateComment) (entity.Comment, error)
	DeleteCommentByID(commentID int64) error
	UpdateCommentByID(ctx context.Context, data types.UpdateComment) (entity.Comment, error)
}
