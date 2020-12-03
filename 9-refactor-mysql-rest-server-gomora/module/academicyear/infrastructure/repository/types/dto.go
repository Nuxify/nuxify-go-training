package types

// CreateAcademicYear create repository types for academic year
type CreateAcademicYear struct {
	TenantID    string
	Name        string
	DisplayName string
	IsActive    bool
	CreatedBy   string
	UpdatedBy   string
}

// GetAcademicYear get repository types for academic year
type GetAcademicYear struct {
	TenantID string
}

// UpdateAcademicYear update repository types for academic year
type UpdateAcademicYear struct {
	ID          int64
	TenantID    string
	Name        string
	DisplayName string
	IsActive    bool
	UpdatedBy   string
}
