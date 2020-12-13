package service

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
	"rest-server/module/discussion/infrastructure/service/types"
)

// DiscussionCommandService handles the post command service logic
type DiscussionCommandService struct {
	repository.DiscussionCommandRepositoryInterface
}

// =================================POST=================================

// CreatePost creates a resource and persist it in repository
func (service *DiscussionCommandService) CreatePost(ctx context.Context, data types.CreatePost) (entity.Post, error) {
	var post repositoryTypes.CreatePost

	post.AuthorID = data.AuthorID
	post.Content = data.Content

	res, err := service.DiscussionCommandRepositoryInterface.InsertPost(post)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeletePostByID delete post by post id
func (service *DiscussionCommandService) DeletePostByID(postID int64) error {
	err := service.DiscussionCommandRepositoryInterface.DeletePostByID(postID)
	if err != nil {
		return err
	}

	return nil
}

// UpdatePostByID updates the resource and persist it in repository
func (service *DiscussionCommandService) UpdatePostByID(ctx context.Context, data types.UpdatePost) (entity.Post, error) {
	var post repositoryTypes.UpdatePost

	post.ID = data.ID
	post.AuthorID = data.AuthorID
	post.Content = data.Content

	res, err := service.DiscussionCommandRepositoryInterface.UpdatePostByID(post)
	if err != nil {
		return res, err
	}

	return res, nil
}

// =================================COMMENT=================================

// CreateComment creates a resource and persist it in repository
func (service *DiscussionCommandService) CreateComment(ctx context.Context, data types.CreateComment) (entity.Comment, error) {
	var comment repositoryTypes.CreateComment

	comment.PostID = data.PostID
	comment.AuthorID = data.AuthorID
	comment.Content = data.Content

	res, err := service.DiscussionCommandRepositoryInterface.InsertComment(comment)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeleteCommentByID delete Comment by comment id
func (service *DiscussionCommandService) DeleteCommentByID(commentID int64) error {
	err := service.DiscussionCommandRepositoryInterface.DeleteCommentByID(commentID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCommentByID updates the resource and persist it in repository
func (service *DiscussionCommandService) UpdateCommentByID(ctx context.Context, data types.UpdateComment) (entity.Comment, error) {
	var comment repositoryTypes.UpdateComment

	comment.ID = data.ID
	comment.Content = data.Content

	res, err := service.DiscussionCommandRepositoryInterface.UpdateCommentByID(comment)
	if err != nil {
		return res, err
	}

	return res, nil
}
