package dao

import (
	"sigmadoc/model"
)

const createDocumentFile string = "/sql/create_doc.sql"

// CreateDocument :
func CreateDocument(doc model.Document) {
	db := OpenDatabase()
	defer db.Close()

	// sql := PrepareQuery(createDocumentFile)

}
