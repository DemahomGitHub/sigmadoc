package service

import (
	"sigmadoc/dao"
	"sigmadoc/model"
)

// FindAllDocuments retreives all documents from the DB
func FindAllDocuments() []model.Document {
	return dao.FindAllDocuments()
}
