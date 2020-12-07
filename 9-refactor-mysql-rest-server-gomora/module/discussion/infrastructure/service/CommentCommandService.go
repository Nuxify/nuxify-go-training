package service

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
	"rest-server/module/discussion/infrastructure/service/types"
)

// CommentCommandService handles the comment command service logic
type CommentCommandService struct {
	repository.CommentCommandRepositoryInterface
}

// CreateComment creates a resource and persist it in repository
func (service *CommentCommandService) CreateComment(ctx context.Context, data types.CreateComment) (entity.Comment, error) {
	var comment repositoryTypes.CreateComment

	comment.PostID = data.PostID
	comment.AuthorID = data.AuthorID
	comment.Content = data.Content

	res, err := service.CommentCommandRepositoryInterface.InsertComment(comment)
	if err != nil {
		return res, err
	}

	return res, nil
}

// DeleteCommentByID delete Comment by comment id
func (service *CommentCommandService) DeleteCommentByID(commentID int64) error {
	err := service.CommentCommandRepositoryInterface.DeleteCommentByID(commentID)
	if err != nil {
		return err
	}

	return nil
}

// UpdateCommentByID updates the resource and persist it in repository
func (service *CommentCommandService) UpdateCommentByID(ctx context.Context, data types.UpdateComment) (entity.Comment, error) {
	var comment repositoryTypes.UpdateComment

	comment.AuthorID = data.AuthorID
	comment.Content = data.Content

	res, err := service.CommentCommandRepositoryInterface.UpdateCommentByID(comment)
	if err != nil {
		return res, err
	}

	return res, nil
}
