package types

// CreateUser create repository types for academic year
type CreateUser struct {
	ID            int64
	Email         string
	FirstName     string
	LastName      string
	ContactNumber string
}

// GetUser get repository types for academic year
type GetUser struct {
	ID int64
}

// UpdateUser update repository types for academic year
type UpdateUser struct {
	ID            int64
	Email         string
	FirstName     string
	LastName      string
	ContactNumber string
}
