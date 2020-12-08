package service

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
	"rest-server/module/discussion/infrastructure/service/types"
)

// PostCommandService handles the post command service logic
type PostCommandService struct {
	repository.PostCommandRepositoryInterface
}

// CreatePost creates a resource and persist it in repository
func (service *PostCommandService) CreatePost(ctx context.Context, data types.CreatePost) (entity.Post, error) {
	var post repositoryTypes.CreatePost

	post.AuthorID = data.AuthorID
	post.Content = data.Content

	res, err := service.PostCommandRepositoryInterface.InsertPost(post)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeletePostByID delete post by post id
func (service *PostCommandService) DeletePostByID(postID int64) error {
	err := service.PostCommandRepositoryInterface.DeletePostByID(postID)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePostByID updates the resource and persist it in repository
func (service *PostCommandService) UpdatePostByID(ctx context.Context, data types.UpdatePost) (entity.Post, error) {
	var post repositoryTypes.UpdatePost

	post.ID = data.ID
	post.AuthorID = data.AuthorID
	post.Content = data.Content

	res, err := service.PostCommandRepositoryInterface.UpdatePostByID(post)
	if err != nil {
		return res, err
	}

	return res, nil
}
