package routes

import (
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
	routes = append(routes, authRoute)

	for _, route := range routes {
		r.HandleFunc(route.URI, route.Function).Methods(route.Method)
	}

	return r
}
