package routes

import (
	"api/src/controllers"
	"net/http"
)

var routesUsers = []Route{
	{
		Uri:         "/users",
		Method:      http.MethodPost,
		Function:    controllers.CreateUser,
		RequireAuth: false,
	}, {
		Uri:         "/users",
		Method:      http.MethodGet,
		Function:    controllers.GetUsers,
		RequireAuth: false,
	}, {
		Uri:         "/users/{userId}",
		Method:      http.MethodGet,
		Function:    controllers.GetOneUser,
		RequireAuth: false,
	}, {
		Uri:         "/users/{userId}",
		Method:      http.MethodPut,
		Function:    controllers.UpdateUser,
		RequireAuth: false,
	}, {
		Uri:         "/users/{userId}",
		Method:      http.MethodDelete,
		Function:    controllers.DeleteUser,
		RequireAuth: false,
	},
}
