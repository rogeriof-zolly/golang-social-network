package controllers

import (
	"devbook/src/responses"
	"io/ioutil"
	"net/http"
)

func Authenticate(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.Err(w, http.StatusUnprocessableEntity, []error{err})
	}

	if requestBody != nil {

	}
}
