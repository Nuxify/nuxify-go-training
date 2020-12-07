package service

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
	"rest-server/module/discussion/infrastructure/service/types"
)

// PostQueryService handles business logic in the service layer
type PostQueryService struct {
	repository.PostQueryRepositoryInterface
}

// GetPosts returns the Posts
func (service *PostQueryService) GetPosts(ctx context.Context, data types.GetPost) ([]entity.Post, error) {
	var post repositoryTypes.GetPost

	post.ID = data.ID

	res, err := service.PostQueryRepositoryInterface.SelectPosts(post)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetPostByID returns the post by id
func (service *PostQueryService) GetPostByID(ctx context.Context, data types.GetPost) ([]entity.Post, error) {
	var post repositoryTypes.GetPost

	post.ID = data.ID

	res, err := service.PostQueryRepositoryInterface.SelectPostByID(post)
	if err != nil {
		return res, err
	}

	return res, nil
}
