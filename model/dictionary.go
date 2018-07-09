package model

import (
	"net/http"
)

// Dictionary represents a table in the databse
type Dictionary struct {
	ID             int    `json:"id"`
	Term           string `json:"term"`
	Frequency      int    `json:"frequency"`
	DocsReferences string `json:"docs_references"`
}

// FindTerm a term the Database
func FindTerm(w http.ResponseWriter, r *http.Request) {}
