package controllers

import "net/http"

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Creating user"))
}

// GetUsers returns all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting all users"))
}

// GetOneUser returns one users from the database by ID
func GetOneUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Getting one user"))
}

// GetOneUser updates one users from the database by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Updating one user"))
}

// GetOneUser removes one users from the database by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Deleting one user"))
}
