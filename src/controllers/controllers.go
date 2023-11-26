package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatal(err)
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		log.Fatal(err)
	}

	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}

	userRepository := repositories.NewUsersRepository(db)
	userRepository.Create(user)
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
