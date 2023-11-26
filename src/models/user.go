package models

import (
	"errors"
	"strings"
	"time"
)

// User is the database model for a user
type User struct {
	ID        uint64    `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Nickname  string    `json:"nickname,omitempty"`
	Email     string    `json:"email,omitempty"`
	Password  string    `json:"password,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
}

func (user *User) Prepare() []error {
	if err := user.validate(); err != nil {
		return err
	}

	user.format()
	return nil
}

func (user *User) validate() []error {
	errorArray := []error{}

	if user.Name == "" {
		errorArray = append(errorArray, errors.New("name should not be empty"))
	}

	if user.Nickname == "" {
		errorArray = append(errorArray, errors.New("nickname should not be empty"))
	}

	if user.Email == "" {
		errorArray = append(errorArray, errors.New("email should not be empty"))
	}

	if user.Password == "" {
		errorArray = append(errorArray, errors.New("password should not be empty"))
	}

	if len(errorArray) > 0 {
		return errorArray
	}

	return nil
}

func (user *User) format() {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)
}
