package repository

import (
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// DiscussionCommandRepositoryInterface holds the implementable methods for the discussion command repository
type DiscussionCommandRepositoryInterface interface {
	DeletePostByID(PostID int64) error
	DeleteCommentByID(CommentID int64) error
	InsertPost(data repositoryTypes.CreatePost) (entity.Post, error)
	InsertComment(data repositoryTypes.CreateComment) (entity.Comment, error)
	UpdatePostByID(data repositoryTypes.UpdatePost) (entity.Post, error)
	UpdateCommentByID(data repositoryTypes.UpdateComment) (entity.Comment, error)
}
