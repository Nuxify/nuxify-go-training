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

// CommentQueryService handles business logic in the service layer
type CommentQueryService struct {
	repository.CommentQueryRepositoryInterface
}

// ====================================POST====================================

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

// ====================================COMMENT====================================

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
