package routes

import (
	"net/http"

	"github.com/tawfeeq0/go_ssl_server/api/controllers"
)

var usersRoutes = []Route{
	Route{
		Uri:     "/users",
		Method:  http.MethodGet,
		Handler: controllers.GetUsers,
	},
	Route{
		Uri:     "/users",
		Method:  http.MethodPost,
		Handler: controllers.CreateUser,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodPut,
		Handler: controllers.UpdateUser,
	},
	Route{
		Uri:     "/users/{id}",
		Method:  http.MethodDelete,
		Handler: controllers.DeleteUser,
	},
	Route{
		Uri:     "/user/{id}",
		Method:  http.MethodGet,
		Handler: controllers.GetUser,
	},
}
