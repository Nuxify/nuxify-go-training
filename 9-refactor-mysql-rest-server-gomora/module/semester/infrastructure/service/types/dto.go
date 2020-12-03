package types

// CreateSemester service struct for create semester
type CreateSemester struct {
	TenantID    string
	Name        string
	DisplayName string
	IsActive    bool
	CreatedBy   string
	UpdatedBy   string
}

// GetSemester service struct for get semester
type GetSemester struct {
	TenantID string
}

// UpdateSemester service struct for update semester
type UpdateSemester struct {
	ID          int64
	TenantID    string
	Name        string
	DisplayName string
	IsActive    bool
	UpdatedBy   string
}
