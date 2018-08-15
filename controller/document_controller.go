package controller

import (
	"log"
	"net/http"
)

// CreateDocument retreives all documents from the DB
func CreateDocument(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		log.Println(err)
	}
}
