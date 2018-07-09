package service

import (
	"sigmadoc/dao"
	"sigmadoc/model"
)

// GetUser gets a user from the DB
func GetUser(username string, password string) model.User {
	return dao.GetUser(username, password)
}

// Authenticate get
func Authenticate(username string, password string) model.User {
	user := model.User{
		Username: "Dem",
		Password: "xxxxx",
		Email:    "demahom@gmail.com",
	}
	return user
}
