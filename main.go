package main

import (
	"log"
	"net/http"
	"sigmadoc/model"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/user", model.GetUser).Methods("GET")

	err := http.ListenAndServe(":3000", router)
	if err != nil {
		log.Fatal(err)
	}
}
