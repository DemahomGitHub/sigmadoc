package model

// Document is a document provided by an administration
// It represents an entity in the database
type Document struct {
	// ID helps to identify a single document in the database
	ID int `json:"id"`
	// Name is the of organization whom provided the document
	Name string `json:"name"`
	// FileLocation is the location of the file
	FileLocation string `json:"file_location"`
	// It is the description of the document
	Description string `json:"description"`
}
