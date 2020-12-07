package repository

import (
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// PostQueryRepositoryInterface holds the methods for the academic year query repository
type PostQueryRepositoryInterface interface {
	SelectPosts(data repositoryTypes.GetPost) ([]entity.Post, error)
	SelectPostByID(data repositoryTypes.GetPost) ([]entity.Post, error)
}
