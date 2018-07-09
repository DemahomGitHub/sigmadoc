package main

import (
	"log"
	"net/http"
	"sigmadoc/controller"
	"sigmadoc/model"

	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	newRouter()
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func newRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/user", controller.GetUser).Methods("GET")
	router.HandleFunc("/documents", model.GetDocument).Methods("GET")
}
