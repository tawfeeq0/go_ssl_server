package server

import (
	"github.com/gorilla/mux"
	"github.com/tawfeeq0/go_ssl_server/server/routes"
)

func NewRouter() *mux.Router {
	r := mux.NewRouter().StrictSlash(true)
	//return routes.SetupRoutes(r)
	return routes.SetupRoutesWithMiddlewares(r)
}
