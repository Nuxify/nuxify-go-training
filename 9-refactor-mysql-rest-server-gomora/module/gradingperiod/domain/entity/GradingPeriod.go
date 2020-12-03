package entity

import (
	"time"
)

// GradingPeriod holds the grading period entity fields
type GradingPeriod struct {
	ID          int
	Name        string
	DisplayName string    `db:"display_name"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// GetModelName returns the model name of grading period entity that can be used for naming schemas
func (entity *GradingPeriod) GetModelName() string {
	return "grading_periods"
}
