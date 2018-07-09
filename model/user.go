package model

// User represente an entity of USER in the database
type User struct {
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}
