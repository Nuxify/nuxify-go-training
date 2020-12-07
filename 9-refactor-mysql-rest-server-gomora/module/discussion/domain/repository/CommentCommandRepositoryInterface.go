package repository

import (
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// CommentCommandRepositoryInterface holds the implementable methods for the comment command repository
type CommentCommandRepositoryInterface interface {
	DeleteCommentByID(CommentID int64) error
	InsertComment(data repositoryTypes.CreateComment) (entity.Comment, error)
	UpdateCommentByID(data repositoryTypes.UpdateComment) (entity.Comment, error)
}
