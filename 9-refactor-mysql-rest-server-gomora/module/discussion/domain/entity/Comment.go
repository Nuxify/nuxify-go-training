package entity

import (
	"time"
)

// Comment holds the comment entity fields
type Comment struct {
	ID        int64
	PostID    int64 `db:"post_id"`
	AuthorID  int64 `db:"author_id"`
	Content   string
	CreatedAt time.Time `db:"created_at"`
	UpdatedAt time.Time `db:"updated_at"`
}

// GetModelName returns the model name of academic year entity that can be used for naming schemas
func (entity *Comment) GetModelName() string {
	return "comments"
}
