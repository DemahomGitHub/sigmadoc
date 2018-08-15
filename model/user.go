package model

import (
	"strings"
)

// User represente an entity of USER in the database
type User struct {
	Email    string `json:"email"`
	Username string `json:"user_name"`
	Password string `json:"password"`
}

// IsEmpty :
func (user User) IsEmpty() bool {
	return !(user.HasEmail() && user.HasPassword() && user.HasUsername())
}

// HasEmail :
func (user User) HasEmail() bool {
	return len(user.Email) > 0
}

// HasUsername :
func (user User) HasUsername() bool {
	return len(user.Username) > 0
}

// HasPassword :
func (user User) HasPassword() bool {
	return len(user.Password) > 0
}

// Equals compares the user name and password of 2 objects
func (user User) Equals(u User) bool {
	sameUsername := strings.Compare(user.Username, u.Username) == 0
	samePassword := strings.Compare(user.Password, u.Password) == 0
	return sameUsername && samePassword
}

// FullyEquals compare all the fields of 2 objects
func (user User) FullyEquals(u User) bool {
	sameUsername := strings.Compare(user.Username, u.Username) == 0
	samePassword := strings.Compare(user.Password, u.Password) == 0
	sameEmail := strings.Compare(user.Email, u.Email) == 0
	return sameEmail && sameUsername && samePassword
}

// Clean removes all spaces in each field of a User object
func (user *User) Clean() {
	user.Email = strings.TrimSpace(user.Email)
	user.Password = strings.TrimSpace(user.Password)
	user.Username = strings.TrimSpace(user.Username)
}
