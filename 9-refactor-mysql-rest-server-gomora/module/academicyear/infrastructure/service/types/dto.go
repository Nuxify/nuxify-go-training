package types

// CreateAcademicYear create service type for academic year
type CreateAcademicYear struct {
	TenantID    string
	Name        string
	DisplayName string
	IsActive    bool
	CreatedBy   string
	UpdatedBy   string
}

// GetAcademicYear get service type for academic year
type GetAcademicYear struct {
	TenantID string
}

// UpdateAcademicYear update service type for academic year
type UpdateAcademicYear struct {
	ID          int64
	TenantID    string
	Name        string
	DisplayName string
	IsActive    bool
	UpdatedBy   string
}
