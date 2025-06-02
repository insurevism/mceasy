package dto

import "time"

// ClientCredential represents a client credential
type ClientCredential struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Fullname string `json:"fullname"`
}

// ClientSession represents a client session
type ClientSession struct {
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	Fullname  string    `json:"fullname"`
	ClientKey string    `json:"client_key"`
	ExpiresAt time.Time `json:"expires_at"`
}
