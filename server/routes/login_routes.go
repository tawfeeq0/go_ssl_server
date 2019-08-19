package routes

import (
	"net/http"

	"github.com/tawfeeq0/go_ssl_server/api/controllers"
)

var loginRoutes = []Route{
	Route{
		Uri:     "/login",
		Method:  http.MethodPost,
		Handler: controllers.Login,
		AuthRequired : false,
	},
}
