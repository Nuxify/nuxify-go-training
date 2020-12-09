package repository

import (
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// PostCommandRepositoryInterface holds the implementable methods for the comment command repository
type PostCommandRepositoryInterface interface {
	DeletePostByID(PostID int64) error
	InsertPost(data repositoryTypes.CreatePost) (entity.Post, error)
	UpdatePostByID(data repositoryTypes.UpdatePost) (entity.Post, error)
}

// CommentCommandRepositoryInterface holds the implementable methods for the comment command repository
type CommentCommandRepositoryInterface interface {
	DeleteCommentByID(CommentID int64) error
	InsertComment(data repositoryTypes.CreateComment) (entity.Comment, error)
	UpdateCommentByID(data repositoryTypes.UpdateComment) (entity.Comment, error)
}
