package types

// CreateUser create service type for user
type CreateUser struct {
	Email         string
	FirstName     string
	LastName      string
	ContactNumber string
}

// GetUser get service type for user
type GetUser struct {
	ID int64
}

// UpdateUser update service type for user
type UpdateUser struct {
	ID            int64
	Email         string
	FirstName     string
	LastName      string
	ContactNumber string
}
