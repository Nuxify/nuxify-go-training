package entity

import (
	"time"
)

// Semester holds the semester entity fields
type Semester struct {
	ID          int64
	TenantID    string `db:"tenant_id"`
	Name        string
	DisplayName string    `db:"display_name"`
	IsActive    bool      `db:"is_active"`
	CreatedBy   string    `db:"created_by"`
	UpdatedBy   string    `db:"updated_by"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

// GetModelName returns the model name of semester entity that can be used for naming schemas
func (entity *Semester) GetModelName() string {
	return "semesters"
}
