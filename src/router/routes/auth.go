package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var authRoute = Route{
	URI:          "/auth",
	Method:       http.MethodPost,
	Function:     controllers.Authenticate,
	AuthRequired: false,
}
