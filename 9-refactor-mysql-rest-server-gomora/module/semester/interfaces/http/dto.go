package types

// CreateSemesterRequest request type for create semester
type CreateSemesterRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
}

// CreateSemesterResponse response type for create semester
type CreateSemesterResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// SemesterResponse response type for semester
type SemesterResponse struct {
	ID          int64  `json:"id"`
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
	CreatedAt   int64  `json:"createdAt"`
	UpdatedAt   int64  `json:"updatedAt"`
}

// UpdateSemesterRequest request type for update semester
type UpdateSemesterRequest struct {
	Name        string `json:"name"`
	DisplayName string `json:"displayName"`
	IsActive    bool   `json:"isActive"`
}
