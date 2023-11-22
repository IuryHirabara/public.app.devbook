package routes

import (
	"net/http"

	"app.devbook/src/controller"
)

var publicationsRoutes = []Route{
	{
		URI:          "/publications",
		Method:       http.MethodPost,
		Func:         controller.CreatePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}/like",
		Method:       http.MethodPost,
		Func:         controller.LikePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}/dislike",
		Method:       http.MethodPost,
		Func:         controller.DislikePublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}/edit",
		Method:       http.MethodGet,
		Func:         controller.LoadEditPublicationPage,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodPut,
		Func:         controller.EditPublication,
		RequiresAuth: true,
	},
	{
		URI:          "/publications/{publicationId}",
		Method:       http.MethodDelete,
		Func:         controller.DeletePublication,
		RequiresAuth: true,
	},
}
