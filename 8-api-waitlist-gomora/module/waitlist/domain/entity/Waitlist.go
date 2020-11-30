package entity

import (
	"time"
)

// Waitlist holds the waitlist entity fields
type Waitlist struct {
	Email     string
	CreatedAt time.Time `db:"created_at"`
}

// GetModelName returns the model name of waitlist entity that can be used for naming schemas
func (entity *Waitlist) GetModelName() string {
	return "waitlists"
}
