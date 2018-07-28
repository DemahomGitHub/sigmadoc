package routers

import (
	"sigmadoc/controller"

	"github.com/gorilla/mux"
)

// InitRoutes initialize all the routes
func InitRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/login", controller.Login).Methods("POST")

	return router
}
