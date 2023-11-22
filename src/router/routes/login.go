package routes

import (
	"net/http"

	"app.devbook/src/controller"
)

var loginRoutes = []Route{
	{
		URI:          "/",
		Method:       http.MethodGet,
		Func:         controller.LoadLoginPage,
		RequiresAuth: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodGet,
		Func:         controller.LoadLoginPage,
		RequiresAuth: false,
	},
	{
		URI:          "/login",
		Method:       http.MethodPost,
		Func:         controller.Login,
		RequiresAuth: false,
	},
}
