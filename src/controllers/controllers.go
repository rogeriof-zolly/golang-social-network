package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, []error{err})
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}
	if validationErrors := user.Prepare(); validationErrors != nil {
		responses.Err(w, http.StatusBadRequest, validationErrors)
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, []error{err})
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	user.ID, err = userRepository.Create(user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, []error{err})
		return
	}

	responses.JSON(w, http.StatusCreated, user)
}

// GetUsers returns all users from the database
func GetUsers(w http.ResponseWriter, r *http.Request) {
	searchParameter := r.URL.Query().Get("user")

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, []error{err})
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	users, err := userRepository.Read(searchParameter)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	responses.JSON(w, http.StatusOK, users)
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
