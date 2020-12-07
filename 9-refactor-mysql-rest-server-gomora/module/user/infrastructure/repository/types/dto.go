package types

// CreateUser create repository types for user
type CreateUser struct {
	ID            int64
	Email         string
	FirstName     string
	LastName      string
	ContactNumber string
}

// GetUser get repository types for user
type GetUser struct {
	ID int64
}

// UpdateUser update repository types for user
type UpdateUser struct {
	ID            int64
	Email         string
	FirstName     string
	LastName      string
	ContactNumber string
}
