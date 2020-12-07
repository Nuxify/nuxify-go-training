package service

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
	"rest-server/module/discussion/infrastructure/service/types"
)

// CommentQueryService handles business logic in the service layer
type CommentQueryService struct {
	repository.CommentQueryRepositoryInterface
}

// GetComments returns the Comments
func (service *CommentQueryService) GetComments(ctx context.Context, data types.GetComment) ([]entity.Comment, error) {
	var comment repositoryTypes.GetComment

	comment.ID = data.ID

	res, err := service.CommentQueryRepositoryInterface.SelectComments(comment)
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetCommentByID returns the comment by id
func (service *CommentQueryService) GetCommentByID(ctx context.Context, data types.GetComment) ([]entity.Comment, error) {
	var comment repositoryTypes.GetComment

	comment.ID = data.ID

	res, err := service.CommentQueryRepositoryInterface.SelectCommentByID(comment)
	if err != nil {
		return res, err
	}

	return res, nil
}
