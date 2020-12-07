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

// SelectPosts select a post by post id
func (repository *PostQueryRepository) SelectPosts(data repositoryTypes.GetPost) ([]entity.Post, error) {
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
