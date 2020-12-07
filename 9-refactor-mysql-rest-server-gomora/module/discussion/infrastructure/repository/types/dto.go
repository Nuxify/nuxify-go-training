package types

// CreatePost create repository types for post
type CreatePost struct {
	ID       int64  `json:"id"`
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// GetPost get repository types for post
type GetPost struct {
	ID int64
}

// UpdatePost update repository types for post
type UpdatePost struct {
	ID       int64  `json:"id"`
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// Author response struct for author
type Author struct {
	ID            int64  `json:"id"`
	Email         string `json:"email"`
	FirstName     string `json:"firstName"`
	LastName      string `json:"lastName"`
	ContactNumber string `json:"contactNumber"`
	CreatedAt     int64  `json:"createdAt"`
	UpdatedAt     int64  `json:"updatedAt"`
}
