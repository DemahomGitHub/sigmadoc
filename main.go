package main

import (
	"log"
	"net/http"
	"sigmadoc/controller"

	"github.com/gorilla/mux"
)

var router *mux.Router

func main() {
	initRouter()
	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}

func initRouter() {
	router = mux.NewRouter()
	router.HandleFunc("/login", controller.Login).Methods("POST")
}
