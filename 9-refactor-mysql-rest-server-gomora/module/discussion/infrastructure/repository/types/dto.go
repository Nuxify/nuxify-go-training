package types

// CreatePost create repository types for post
type CreatePost struct {
	ID       int64
	AuthorID int64
	Content  string
}

// GetPost get repository types for post
type GetPost struct {
	ID int64
}

// UpdatePost update repository types for post
type UpdatePost struct {
	ID       int64
	AuthorID int64
	Content  string
}

// CreateComment create repository types for comment
type CreateComment struct {
	ID       int64
	PostID   int64
	AuthorID int64
	Content  string
}

// GetComment get repository types for comment
type GetComment struct {
	ID int64
}

// UpdateComment update repository types for comment
type UpdateComment struct {
	ID       int64
	AuthorID int64
	Content  string
}
