package types

// CreatePost create service type for post
type CreatePost struct {
	AuthorID int64
	Content  string
}

// GetPost get service type for post
type GetPost struct {
	ID int64
}

// UpdatePost update service type for post
type UpdatePost struct {
	ID       int64
	AuthorID int64
	Content  string
}
