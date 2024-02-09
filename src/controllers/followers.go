package controllers

import (
	"database/sql"
	"devbook/src/authentication"
	"devbook/src/database"
	"devbook/src/repositories"
	"devbook/src/responses"
	"errors"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Gets all followers from a given user
func RetrieveAllFollowers(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	followersRepository := repositories.NewFollowersRepository(db)

	followers, err := followersRepository.GetAllFollowers(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
	}

	responses.JSON(w, http.StatusOK, followers)
}

func RetrieveUsersFollowing(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	userID, err := strconv.ParseUint(params["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	followersRepository := repositories.NewFollowersRepository(db)

	following, err := followersRepository.GetFollowing(userID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
	}

	responses.JSON(w, http.StatusOK, following)
}

// Follow a User
func FollowUser(w http.ResponseWriter, r *http.Request) {
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

	if userIdFromToken == userID {
		cannotFollowYourself := errors.New("you cannot follow yourself")
		responses.Err(w, http.StatusBadRequest, []error{cannotFollowYourself})
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}
	defer db.Close()

	usersRepository := repositories.NewUsersRepository(db)

	userToFollow, err := usersRepository.ReadByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("user to follow does not exist")
		}
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	followersRepository := repositories.NewFollowersRepository(db)

	err = followersRepository.FollowUser(userIdFromToken, userToFollow.ID)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}

func UnfollowUser(w http.ResponseWriter, r *http.Request) {
	parameters := mux.Vars(r)

	userID, err := strconv.ParseUint(parameters["userId"], 10, 64)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	userIDFromToken, err := authentication.ExtractUserIdFromToken(r)
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	if userIDFromToken == userID {
		cannotUnfollowYourself := errors.New("cannot unfollow yourself")
		responses.Err(w, http.StatusBadRequest, []error{cannotUnfollowYourself})
		return
	}

	db, err := database.Connect()
	if err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}
	defer db.Close()

	userRepository := repositories.NewUsersRepository(db)

	userToUnfollow, err := userRepository.ReadByID(userID)
	if err != nil {
		if err == sql.ErrNoRows {
			err = errors.New("user to unfollow does not exist")
		}
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	followersRepository := repositories.NewFollowersRepository(db)

	if err = followersRepository.UnfollowUser(
		userIDFromToken, userToUnfollow.ID,
	); err != nil {
		responses.Err(w, http.StatusBadRequest, []error{err})
		return
	}

	responses.JSON(w, http.StatusNoContent, nil)
}
