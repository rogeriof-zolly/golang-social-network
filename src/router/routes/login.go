package routes

import (
	"devbook/src/controllers"
	"net/http"
)

var LoginRoute = Route{
	URI:          "/login",
	Method:       http.MethodPost,
	Function:     controllers.Login,
	AuthRequired: false,
}
