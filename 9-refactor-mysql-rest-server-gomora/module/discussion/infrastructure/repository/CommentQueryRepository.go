package repository

import (
	"errors"
	"fmt"

	"rest-server/infrastructures/database/mysql/types"
	apiError "rest-server/internal/errors"
	"rest-server/module/discussion/domain/entity"
	repositoryTypes "rest-server/module/discussion/infrastructure/repository/types"
)

// CommentQueryRepository handles databas access logic
type CommentQueryRepository struct {
	types.MySQLDBHandlerInterface
}

// SelectComments select a comment by comment id
func (repository *CommentQueryRepository) SelectComments(data repositoryTypes.GetComment) ([]entity.Comment, error) {
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
