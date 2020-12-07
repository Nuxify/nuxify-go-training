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

// CreateComment create service type for Comment
type CreateComment struct {
	PostID   int64
	AuthorID int64
	Content  string
}

// GetComment get service type for Comment
type GetComment struct {
	ID int64
}

// UpdateComment update service type for Comment
type UpdateComment struct {
	ID       int64
	AuthorID int64
	Content  string
}
