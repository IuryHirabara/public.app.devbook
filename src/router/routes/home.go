package routes

import (
	"net/http"

	"app.devbook/src/controller"
)

var homeRoute = Route{
	URI:          "/home",
	Method:       http.MethodGet,
	Func:         controller.LoadHomePage,
	RequiresAuth: true,
}
