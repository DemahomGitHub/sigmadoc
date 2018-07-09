package controller

import (
	"encoding/json"
	"net/http"
	"sigmadoc/service"
)

// GetUser gets an entity of user from the USER table
func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(service.GetUser("demahom", "1234"))
}

// Authenticate allow user to authenticate
func Authenticate(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(service.Authenticate("Dem", "xxxx"))
}
