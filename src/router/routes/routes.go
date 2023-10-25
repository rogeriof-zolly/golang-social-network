package routes

import "net/http"

// Route defines the structure of the API routes
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}
