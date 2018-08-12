package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"sigmadoc/model"
	"sigmadoc/service"
)

// Login allows a user to authenticate
func Login(w http.ResponseWriter, r *http.Request) {
	perr := r.ParseForm()

	if perr != nil {
		log.Fatal(perr)
	}

	form := model.User{Username: r.FormValue("user_name"), Password: r.FormValue("password")}
	resp := service.GetUserByUsernameAndPassword(form)

	if resp.Status == model.FAILED {
		json.NewEncoder(w).Encode(resp.Message)
	} else {
		json.NewEncoder(w).Encode(resp.UserResp)
	}
}

// CreateUser adds a new user in the DB
func CreateUser(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()

	if err != nil {
		log.Fatal(err)
	}

	form := model.User{
		Username: r.FormValue("user_name"),
		Password: r.FormValue("password"),
		Email:    r.FormValue("email"),
	}

	resp := service.CreateUser(form)

	if resp.Status == model.FAILED {
		json.NewEncoder(w).Encode(resp.Message)
	} else {
		json.NewEncoder(w).Encode(resp.Message)
	}
}
