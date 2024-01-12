package controllers

import (
	"devbook/src/database"
	"devbook/src/models"
	"devbook/src/repositories"
	"devbook/src/responses"
	"encoding/json"
	"io"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	var user models.User
	if err = json.Unmarshal(body, &user); err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)
	userInDatabase, err := userRepository.ReadByEmail(user.Email)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
	}

}
