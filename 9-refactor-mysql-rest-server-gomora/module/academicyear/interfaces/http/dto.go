package types

// CreateAcademicYearRequest request type for create academic year
type CreateAcademicYearRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
}

// CreateAcademicYearResponse response type for create academic year
type CreateAcademicYearResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// AcademicYearResponse response type for academic year
type AcademicYearResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// UpdateAcademicYearRequest request type for update academic year
type UpdateAcademicYearRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
}
