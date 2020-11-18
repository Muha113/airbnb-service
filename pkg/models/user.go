package models

import "time"

// User contains all needful information about user
type User struct {
	ID          int            `json:"id"`
	FirstName   string         `json:"first_name"`
	LastName    string         `json:"last_name"`
	Birthday    time.Time      `json:"birthday"`
	Email       string         `json:"email"`
	Password    string         `json:"password"`
	PhoneNumber string         `json:"phone_number"`
	Username    string         `json:"username"`
	Role        string         `json:"role"`
	Lodgings    []*Lodging     `json:"lodging"`
	UserRequest []*UserRequest `json:"user_request"`
	StarsAmount int            `json:"stars_amount"`
	VotedAmount int            `json:"voted_amount"`
}
