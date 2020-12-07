package application

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/infrastructure/service/types"
)

// CommentCommandServiceInterface holds the implementable method for the comment command service
type CommentCommandServiceInterface interface {
	CreateComment(ctx context.Context, data types.CreateComment) (entity.Comment, error)
	DeleteCommentByID(commentID int64) error
	UpdateCommentByID(ctx context.Context, data types.UpdateComment) (entity.Comment, error)
}
