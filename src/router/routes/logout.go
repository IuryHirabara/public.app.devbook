package routes

import (
	"net/http"

	"app.devbook/src/controller"
)

var logoutRoute = Route{
	URI:          "/logout",
	Method:       http.MethodGet,
	Func:         controller.Logout,
	RequiresAuth: true,
}
