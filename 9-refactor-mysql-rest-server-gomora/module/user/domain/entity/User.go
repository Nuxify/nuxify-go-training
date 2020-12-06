package entity

import (
	"time"
)

// User holds the user entity fields
type User struct {
	ID            int64
	Email         string
	FirstName     string    `db:"first_name"`
	LastName      string    `db:"last_name"`
	ContactNumber string    `db:"contact_number"`
	CreatedAt     time.Time `db:"created_at"`
	UpdatedAt     time.Time `db:"updated_at"`
}

// GetModelName returns the model name of academic year entity that can be used for naming schemas
func (entity *User) GetModelName() string {
	return "users"
}
