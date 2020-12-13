package service

import (
	"context"

	"rest-server/module/discussion/domain/entity"
	"rest-server/module/discussion/domain/repository"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
	"rest-server/module/discussion/infrastructure/service/types"
)

// DiscussionQueryService handles business logic in the service layer
type DiscussionQueryService struct {
	repository.DiscussionQueryRepositoryInterface
}

// ====================================POST====================================

// GetPosts returns the Posts
func (service *DiscussionQueryService) GetPosts(ctx context.Context) ([]entity.Post, error) {
	res, err := service.DiscussionQueryRepositoryInterface.SelectPosts()
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetPostByID returns the post by id
func (service *DiscussionQueryService) GetPostByID(ctx context.Context, data types.GetPost) ([]entity.Post, error) {
	var post repositoryTypes.GetPost

	post.ID = data.ID

	res, err := service.DiscussionQueryRepositoryInterface.SelectPostByID(post)
	if err != nil {
		return res, err
	}

	return res, nil
}

// ====================================COMMENT====================================

// GetComments returns the Comments
func (service *DiscussionQueryService) GetComments(ctx context.Context) ([]entity.Comment, error) {
	res, err := service.DiscussionQueryRepositoryInterface.SelectComments()
	if err != nil {
		return res, err
	}

	return res, nil
}

// GetCommentByID returns the comment by id
func (service *DiscussionQueryService) GetCommentByID(ctx context.Context, data types.GetComment) ([]entity.Comment, error) {
	var comment repositoryTypes.GetComment

	comment.ID = data.ID

	res, err := service.DiscussionQueryRepositoryInterface.SelectCommentByID(comment)
	if err != nil {
		return res, err
	}

	return res, nil
}
