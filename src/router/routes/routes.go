package routes

import (
	"net/http"

	"app.devbook/src/middlewares"
	"github.com/gorilla/mux"
)

// Route representa todas as rotas do app
type Route struct {
	URI          string
	Method       string
	Func         func(http.ResponseWriter, *http.Request)
	RequiresAuth bool
}

// Configure coloca todas as rotas dentro do router
func Configure(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, homeRoute)
	routes = append(routes, publicationsRoutes...)
	routes = append(routes, logoutRoute)

	for _, route := range routes {
		if route.RequiresAuth {
			router.HandleFunc(route.URI, middlewares.Logger(
				middlewares.Auth(route.Func),
			)).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI, middlewares.Logger(route.Func)).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}
