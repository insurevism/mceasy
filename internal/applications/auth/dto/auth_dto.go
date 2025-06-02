package dto

import "time"

// ClientCredential represents a client credential
type ClientCredential struct {
	ID       uint64 `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ClientSession represents a client session
type ClientSession struct {
	ID        uint64    `json:"id"`
	ClientKey string    `json:"client_key"`
	ExpiresAt time.Time `json:"expires_at"`
}
