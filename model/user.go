package model

import (
	"fmt"
	"net/http"
)

// User represente an entity of USER in the database
type User struct {
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

// GetUser gets an entity of user from the USER table
func GetUser(w http.ResponseWriter, r *http.Request) {
	user := User{Email: "demahom@mail.com", Username: "Dem", Password: "xxxxx"}
	fmt.Println(w, user)
}
