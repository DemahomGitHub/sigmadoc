package service

import (
	"sigmadoc/dao"
	"sigmadoc/model"
)

// CreateDocument retreives all documents from the DB
func CreateDocument(doc model.Document) {
	dao.CreateDocument(doc)
}
