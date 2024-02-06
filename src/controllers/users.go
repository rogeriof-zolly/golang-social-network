package controllers

import (
	"devbook/src/authentication"
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// CreateUser creates a new user in the database
func CreateUser(w http.ResponseWriter, r *http.Request) {
	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, []error{err})
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}
	// Method passed to Prepare, because this can prepare for POST and PUT requests
	if validationErrors := user.Prepare(r.Method); validationErrors != nil {
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
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, []error{err})
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	user, err := userRepository.ReadByID(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	responses.JSON(w, http.StatusOK, user)
}

// GetOneUser updates one users from the database by ID
func UpdateUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	userIdFromToken, err := authentication.ExtractUserIdFromToken(r)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	if userIdFromToken != userID {
		cannotUpdateError := errors.New("cannot update a user that is not yours")
		responses.Err(w, http.StatusForbidden, []error{cannotUpdateError})
		return
	}

	requestBody, err := io.ReadAll(r.Body)

	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, []error{err})
		return
	}

	var user models.User
	if err = json.Unmarshal(requestBody, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}
	// Method passed to Prepare, because this can prepare for POST and PUT requests
	if validationErrors := user.Prepare(r.Method); validationErrors != nil {
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
	err = userRepository.Update(userID, user)
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, []error{err})
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

// GetOneUser removes one users from the database by ID
func DeleteUser(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	userIdFromToken, err := authentication.ExtractUserIdFromToken(r)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	if userIdFromToken != userID {
		cannotDeleteUser := errors.New("cannot delete a user that is not yours")
		responses.Err(w, http.StatusForbidden, []error{cannotDeleteUser})
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusInternalServerError, []error{err})
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	err = userRepository.Delete(userID)

	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
	}

	responses.JSON(w, http.StatusOK, nil)
}
