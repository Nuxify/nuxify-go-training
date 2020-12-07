package application

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/infrastructure/service/types"
)

// CommentQueryServiceInterface holds the implementable method for the comment query service
type CommentQueryServiceInterface interface {
	GetComments(ctx context.Context, data types.GetComment) ([]entity.Comment, error)
	GetCommentByID(ctx context.Context, data types.GetComment) ([]entity.Comment, error)
}
