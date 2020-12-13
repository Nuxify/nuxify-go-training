package repository

import (
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// DiscussionQueryRepositoryInterface holds the methods for the discussion query repository
type DiscussionQueryRepositoryInterface interface {
	SelectPosts() ([]entity.Post, error)
	SelectPostByID(data repositoryTypes.GetPost) ([]entity.Post, error)
	SelectComments() ([]entity.Comment, error)
	SelectCommentByID(data repositoryTypes.GetComment) ([]entity.Comment, error)
}
