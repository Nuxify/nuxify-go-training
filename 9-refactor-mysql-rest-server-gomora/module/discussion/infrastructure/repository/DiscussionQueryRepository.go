package repository

import (
	"errors"
	"fmt"

	"rest-server/infrastructures/database/mysql/types"
	apiError "rest-server/internal/errors"
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// PostQueryRepository handles databas access logic
type PostQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// CommentQueryRepository handles databas access logic
type CommentQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// =========================================POST=========================================

// SelectPosts select a posts
func (repository *PostQueryRepository) SelectPosts() ([]entity.Post, error) {
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
func (repository *PostQueryRepository) SelectPostByID(data repositoryTypes.GetPost) ([]entity.Post, error) {
	var post entity.Post
	var posts []entity.Post

	stmt := fmt.Sprintf("SELECT * FROM %s Where id=:id", post.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{"id": data.ID}, &posts)
	if err != nil {
		return posts, errors.New(apiError.DatabaseError)
	} else if len(posts) == 0 {
		return posts, errors.New(apiError.MissingRecord)
	}

	return posts, nil
}

// =========================================COMMENT=========================================

// SelectComments select a comments
func (repository *CommentQueryRepository) SelectComments() ([]entity.Comment, error) {
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
func (repository *CommentQueryRepository) SelectCommentByID(data repositoryTypes.GetComment) ([]entity.Comment, error) {
	var comment entity.Comment
	var comments []entity.Comment

	stmt := fmt.Sprintf("SELECT * FROM %s Where id=:id", comment.GetModelName())

	err := repository.Query(stmt, map[string]interface{}{"id": data.ID}, &comments)
	if err != nil {
		return comments, errors.New(apiError.DatabaseError)
	} else if len(comments) == 0 {
		return comments, errors.New(apiError.MissingRecord)
	}

	return comments, nil
}
