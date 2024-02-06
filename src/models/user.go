package models

import (
	"devbook/src/security"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/badoux/checkmail"
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

func (user *User) Prepare(httpMethod string) []error {
	if err := user.validate(httpMethod); err != nil {
		return err
	}

	if err := user.format(httpMethod); err != nil {
		return err
	}

	return nil
}

func (user *User) validate(httpMethod string) []error {
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

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		errorArray = append(errorArray, errors.New("invalid e-mail address"))
	}

	if httpMethod == http.MethodPost && user.Password == "" {
		errorArray = append(errorArray, errors.New("password should not be empty"))
	}

	if len(errorArray) > 0 {
		return errorArray
	}

	return nil
}

func (user *User) format(httpMethod string) []error {
	user.Name = strings.TrimSpace(user.Name)
	user.Nickname = strings.TrimSpace(user.Nickname)
	user.Email = strings.TrimSpace(user.Email)

	if httpMethod == http.MethodPost {
		hashedPassword, err := security.Hash(user.Password)
		if err != nil {
			return []error{err}
		}
		user.Password = string(hashedPassword)
	}

	return nil
}
