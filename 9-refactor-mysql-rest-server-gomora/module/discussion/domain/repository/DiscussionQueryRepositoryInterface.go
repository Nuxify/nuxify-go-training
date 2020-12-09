package repository

import (
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// PostQueryRepositoryInterface holds the methods for the comment query repository
type PostQueryRepositoryInterface interface {
	SelectPosts(data repositoryTypes.GetPost) ([]entity.Post, error)
	SelectPostByID(data repositoryTypes.GetPost) ([]entity.Post, error)
}

// CommentQueryRepositoryInterface holds the methods for the comment query repository
type CommentQueryRepositoryInterface interface {
	SelectComments(data repositoryTypes.GetComment) ([]entity.Comment, error)
	SelectCommentByID(data repositoryTypes.GetComment) ([]entity.Comment, error)
}
