package models

import "time"

// User is the database model for a user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"passoword,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}
