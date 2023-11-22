package routes

import (
	"net/http"

	"app.devbook/src/controller"
)

var userRoutes = []Route{
	{
		URI:          "/cadastro",
		Method:       http.MethodGet,
		Func:         controller.LoadRegisterPage,
		RequiresAuth: false,
	},
	{
		URI:          "/users",
		Method:       http.MethodPost,
		Func:         controller.CreateUser,
		RequiresAuth: false,
	},
	{
		URI:          "/searchUsers",
		Method:       http.MethodGet,
		Func:         controller.LoadUsersPage,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}",
		Method:       http.MethodGet,
		Func:         controller.LoadUserProfile,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/unfollow",
		Method:       http.MethodPost,
		Func:         controller.UnfollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/users/{userId}/follow",
		Method:       http.MethodPost,
		Func:         controller.FollowUser,
		RequiresAuth: true,
	},
	{
		URI:          "/profile",
		Method:       http.MethodGet,
		Func:         controller.LoadLoggedUserProfile,
		RequiresAuth: true,
	},
	{
		URI:          "/edit-user",
		Method:       http.MethodGet,
		Func:         controller.LoadUserEditPage,
		RequiresAuth: true,
	},
	{
		URI:          "/edit-user",
		Method:       http.MethodPut,
		Func:         controller.EditUser,
		RequiresAuth: true,
	},
	{
		URI:          "/update-password",
		Method:       http.MethodGet,
		Func:         controller.LoadUpdatePasswordPage,
		RequiresAuth: true,
	},
	{
		URI:          "/update-password",
		Method:       http.MethodPost,
		Func:         controller.UpdatePassword,
		RequiresAuth: true,
	},
	{
		URI:          "/delete-user",
		Method:       http.MethodDelete,
		Func:         controller.DeleteUser,
		RequiresAuth: true,
	},
}
