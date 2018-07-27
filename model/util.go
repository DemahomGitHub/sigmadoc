package model

import (
	"strings"
)

// Response is a response from the DB
type Response interface {
	HasErrors() bool
	ToJSON() string
}

// UserResponse :
type UserResponse struct {
	Status int
	Data   User
}

// DocumentResponse :
type DocumentResponse struct {
	Status int
	Data   Document
}

// ErrorResponse :
type ErrorResponse struct {
	Status int
	Info   string
}

// Str is string redefined
type Str string

// Equals compare 2 strings and returns 0 if they are equals
func (a Str) Equals(b string) bool {
	return strings.Compare(string(a), b) == 0
}
