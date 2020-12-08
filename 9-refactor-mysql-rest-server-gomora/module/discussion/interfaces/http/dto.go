package types

// CreatePostRequest request type for create post
type CreatePostRequest struct {
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// CreatePostResponse response type for create post
type CreatePostResponse struct {
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// PostResponse response type for post
type PostResponse struct {
	ID        int64  `json:"id"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// UpdatePostRequest request type for update post
type UpdatePostRequest struct {
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// UpdatePostResponse response type for update post
type UpdatePostResponse struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
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

// CreateCommentRequest request type for create Comment
type CreateCommentRequest struct {
	PostID   int64  `json:"postId"`
	AuthorID int64  `json:"authorId"`
	Content  string `json:"content"`
}

// CreateCommentResponse response type for create comment
type CreateCommentResponse struct {
	PostID    int64  `json:"postId"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// CommentResponse response type for comment
type CommentResponse struct {
	ID        int64  `json:"id"`
	PostID    int64  `json:"postId"`
	Content   string `json:"content"`
	CreatedAt int64  `json:"createdAt"`
	UpdatedAt int64  `json:"updatedAt"`
}

// UpdateCommentRequest request type for update comment
type UpdateCommentRequest struct {
	Content string `json:"content"`
}

// UpdateCommentResponse response type for comment
type UpdateCommentResponse struct {
	ID      int64  `json:"id"`
	Content string `json:"content"`
}
