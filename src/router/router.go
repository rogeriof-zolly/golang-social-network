package router

import (
	"devbook/src/router/routes"

	"github.com/gorilla/mux"
)

// Generate returns a router with the configured routes
func Generate() *mux.Router {
	r := mux.NewRouter()

	return routes.SetUp(r)
}
