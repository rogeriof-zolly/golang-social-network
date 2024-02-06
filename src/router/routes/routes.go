package routes

import (
	"devbook/src/middleware"
	"net/http"

	"github.com/gorilla/mux"
)

// Route defines the structure of the API routes
type Route struct {
	URI          string
	Method       string
	Function     func(http.ResponseWriter, *http.Request)
	AuthRequired bool
}

func SetUp(r *mux.Router) *mux.Router {
	routes := userRoutes
	routes = append(routes, LoginRoute)
	routes = append(routes, FollowersRoutes...)

	for _, route := range routes {
		if route.AuthRequired {
			// Defines the function with auth in a variable to make it readable
			authenticatedFunc := middleware.Authenticate(route.Function)

			r.HandleFunc(route.URI, middleware.Logger(authenticatedFunc)).
				Methods(route.Method)
		} else {
			r.HandleFunc(route.URI, middleware.Logger(route.Function)).
				Methods(route.Method)
		}
	}

	return r
}
