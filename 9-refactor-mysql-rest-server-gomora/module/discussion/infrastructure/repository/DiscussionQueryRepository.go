package repository

import (
	"errors"
	"fmt"

	"rest-server/infrastructures/database/mysql/types"
	apiError "rest-server/internal/errors"
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// DiscussionQueryRepository handles databas access logic
type DiscussionQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// =========================================POST=========================================

// SelectPosts select a posts
func (repository *DiscussionQueryRepository) SelectPosts() ([]entity.Post, error) {
	var post entity.Post
	var posts []entity.Post

	stmt := fmt.Sprintf("SELECT * FROM %s", post.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{}, &posts)
	if err != nil {
		return posts, errors.New(apiError.DatabaseError)
	} else if len(posts) == 0 {
		return posts, errors.New(apiError.MissingRecord)
	}

	return posts, nil
}

// SelectPostByID select a post by post id
func (repository *DiscussionQueryRepository) SelectPostByID(data repositoryTypes.GetPost) (entity.Post, error) {
	var post entity.Post
	var posts []entity.Post

	stmt := fmt.Sprintf("SELECT * FROM %s Where id=:id", post.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{"id": data.ID}, &posts)
	if err != nil {
		return entity.Post{}, errors.New(apiError.DatabaseError)
	} else if len(posts) == 0 {
		return entity.Post{}, errors.New(apiError.MissingRecord)
	}

	return posts[0], nil
}

// =========================================COMMENT=========================================

// SelectComments select a comments
func (repository *DiscussionQueryRepository) SelectComments() ([]entity.Comment, error) {
	var comment entity.Comment
	var comments []entity.Comment

	stmt := fmt.Sprintf("SELECT * FROM %s", comment.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{}, &comments)
	if err != nil {
		return comments, errors.New(apiError.DatabaseError)
	} else if len(comments) == 0 {
		return comments, errors.New(apiError.MissingRecord)
	}

	return comments, nil
}

// SelectCommentByID select a comment by comment id
func (repository *DiscussionQueryRepository) SelectCommentByID(data repositoryTypes.GetComment) (entity.Comment, error) {
	var comment entity.Comment
	var comments []entity.Comment

	stmt := fmt.Sprintf("SELECT * FROM %s Where id=:id", comment.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{"id": data.ID}, &comments)
	if err != nil {
		return entity.Comment{}, errors.New(apiError.DatabaseError)
	} else if len(comments) == 0 {
		return entity.Comment{}, errors.New(apiError.MissingRecord)
	}

	return comments[0], nil
}
