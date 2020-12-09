package repository

import (
	"errors"
	"fmt"
	"strings"

	"rest-server/infrastructures/database/mysql/types"
	apiError "rest-server/internal/errors"
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// PostCommandRepository handles the post command repository logic
type PostCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// CommentCommandRepository handles the comment command repository logic
type CommentCommandRepository struct {
	types.MySQLDBHandlerInterface
}

// ==========================================POST==========================================

// DeletePostByID removes post by id
func (repository *PostCommandRepository) DeletePostByID(postID int64) error {
	post := &entity.Post{
		ID: postID,
	}

	// delete post
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", post.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, post)
	if err != nil {
		return err
	}

	return nil
}

// InsertPost creates a new post resource
func (repository *PostCommandRepository) InsertPost(data repositoryTypes.CreatePost) (entity.Post, error) {
	post := &entity.Post{
		AuthorID: data.AuthorID,
		Content:  data.Content,
	}

	// insert post
	stmt := fmt.Sprintf("INSERT INTO %s(author_id, content) VALUES (:author_id, :content)", post.GetModelName())
	res, err := repository.MySQLDBHandlerInterface.Execute(stmt, post)
	if err != nil {
		var errStr string

		if strings.Contains(err.Error(), "Duplicate entry") {
			errStr = apiError.DuplicateRecord
		} else {
			errStr = apiError.DatabaseError
		}

		return *post, errors.New(errStr)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return *post, errors.New(apiError.DatabaseError)
	}

	return *post, nil
}

// UpdatePostByID update resource
func (repository *PostCommandRepository) UpdatePostByID(data repositoryTypes.UpdatePost) (entity.Post, error) {
	post := &entity.Post{
		ID:       data.ID,
		AuthorID: data.AuthorID,
		Content:  data.Content,
	}
	// update post
	stmt := fmt.Sprintf("UPDATE %s SET author_id=:author_id, content=:content WHERE id=:id", post.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, post)
	if err != nil {
		fmt.Println(err)
		return *post, errors.New(apiError.DatabaseError)
	}

	return *post, nil
}

// ==========================================COMMENT==========================================

// DeleteCommentByID removes Comment by id
func (repository *CommentCommandRepository) DeleteCommentByID(commentID int64) error {
	comment := &entity.Comment{
		ID: commentID,
	}

	// delete post
	stmt := fmt.Sprintf("DELETE FROM %s WHERE id=:id", comment.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, comment)
	if err != nil {
		return err
	}

	return nil
}

// InsertComment creates a new comment resource
func (repository *CommentCommandRepository) InsertComment(data repositoryTypes.CreateComment) (entity.Comment, error) {
	comment := &entity.Comment{
		PostID:   data.PostID,
		AuthorID: data.AuthorID,
		Content:  data.Content,
	}

	// insert comment
	stmt := fmt.Sprintf("INSERT INTO %s (post_id, author_id, content) VALUES (:post_id, :author_id, :content)", comment.GetModelName())
	res, err := repository.MySQLDBHandlerInterface.Execute(stmt, comment)
	if err != nil {
		var errStr string

		if strings.Contains(err.Error(), "Duplicate entry") {
			errStr = apiError.DuplicateRecord
		} else {
			errStr = apiError.DatabaseError
		}

		return *comment, errors.New(errStr)
	}
	_, err = res.LastInsertId()
	if err != nil {
		return *comment, errors.New(apiError.DatabaseError)
	}

	return *comment, nil
}

// UpdateCommentByID update resource
func (repository *CommentCommandRepository) UpdateCommentByID(data repositoryTypes.UpdateComment) (entity.Comment, error) {
	comment := &entity.Comment{
		ID:      data.ID,
		Content: data.Content,
	}
	// update comment
	stmt := fmt.Sprintf("UPDATE %s SET content=:content WHERE id=:id", comment.GetModelName())
	_, err := repository.MySQLDBHandlerInterface.Execute(stmt, comment)
	if err != nil {
		fmt.Println(err)
		return *comment, errors.New(apiError.DatabaseError)
	}

	return *comment, nil
}
