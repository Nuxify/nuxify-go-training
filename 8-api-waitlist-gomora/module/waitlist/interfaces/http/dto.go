package http

// CreateWaitlistRequest request struct for create waitlist
type CreateWaitlistRequest struct {
	Email string `json:"email"`
}

// WaitlistResponse response struct
type WaitlistResponse struct {
	Email     string `json:"email"`
	CreatedAt int64  `json:"createdAt"`
}
