package responses

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSON returns a Json response to the request
func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	// No need to return an empty body
	// Status' such as No Content won't allow to return a body
	if data == nil {
		return
	}

	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatal(err)
	}
}

// Err return an error formatted in JSON
func Err(w http.ResponseWriter, statusCode int, errors []error) {
	errorArray := []string{}
	for _, err := range errors {
		errorArray = append(errorArray, err.Error())
	}

	errorsBody := struct {
		Error []string `json:"error"`
	}{
		Error: errorArray,
	}

	JSON(w, statusCode, errorsBody)
}
