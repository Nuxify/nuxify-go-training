package entity

import (
	"time"
)

// Post holds the post entity fields
type Post struct {
	ID        int64
	AuthorID  int64 `db:"author_id"`
	Content   string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetModelName returns the model name of academic year entity that can be used for naming schemas
func (entity *Post) GetModelName() string {
	return "posts"
}
