package models

// UserRequest ...
type UserRequest struct {
	ID       int        `json:"id"`
	User     *User      `json:"user"`
	Requests []*Request `json:"requests"`
	Type     string     `json:"type"`
}

// Request ...
type Request struct {
	ID           int            `json:"id"`
	UserRequests []*UserRequest `json:"user_requests"`
	Status       string         `json:"status"`
}
