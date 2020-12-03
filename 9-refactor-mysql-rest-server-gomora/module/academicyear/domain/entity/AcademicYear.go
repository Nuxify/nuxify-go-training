package entity

import (
	"time"
)

// AcademicYear holds the academic year entity fields
type AcademicYear struct {
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

// GetModelName returns the model name of academic year entity that can be used for naming schemas
func (entity *AcademicYear) GetModelName() string {
	return "academic_years"
}
