package types

// CreatePostRequest request type for create post
type CreatePostRequest struct {
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// CreatePostResponse response type for create post
type CreatePostResponse struct {
	ID        int64  `json:"id"`
	Author    Author `json:"author"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// PostResponse response type for post
type PostResponse struct {
	ID        int64  `json:"id"`
	Author    Author `json:"author"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// UpdatePostRequest request type for update post
type UpdatePostRequest struct {
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
