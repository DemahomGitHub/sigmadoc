package controller

import (
	"encoding/json"
	"net/http"
	"sigmadoc/service"
)

// FindAllDocuments retreives all documents from the DB
func FindAllDocuments(w http.ResponseWriter) {
	json.NewEncoder(w).Encode(service.FindAllDocuments())
}
