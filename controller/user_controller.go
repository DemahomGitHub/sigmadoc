package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sigmadoc/service"
)

// Login allows a user to authenticate
func Login(w http.ResponseWriter, r *http.Request) {
	perr := r.ParseForm()
	if perr != nil {
		log.Fatal(perr)
	}

	resp, err := service.GetUserByUsernameAndPassword(r.FormValue("user_name"), r.FormValue("password"))

	if resp.Data.IsEmpty() {
		json.NewEncoder(w).Encode(err)
	} else {
		json.NewEncoder(w).Encode(resp)
	}
}
