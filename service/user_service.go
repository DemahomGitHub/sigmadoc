package service

import (
	"sigmadoc/dao"
	"sigmadoc/model"

	"golang.org/x/crypto/bcrypt"
)

var responseMessage model.ResponseMessage

// GetUserByUsernameAndPassword gets a user from the DB
func GetUserByUsernameAndPassword(form model.User) model.ResponseMessage {
	// Clean all white spaces
	model.RemoveSpaces(&form)

	if !form.HasUsername() || !form.HasPassword() {
		responseMessage.Status = model.FAILED
		responseMessage.Message = "The User name and the password are required"
		return responseMessage
	}

	userFromDB, err := dao.GetUserByUsernameAndPassword(form)

	if err != nil {
		responseMessage.Status = model.FAILED
		responseMessage.Message = "This user doesn't exist"
		return responseMessage
	}

	isWrongPassword := bcrypt.CompareHashAndPassword([]byte(userFromDB.Password), []byte(form.Password)) != nil

	if isWrongPassword {
		responseMessage.Status = model.FAILED
		responseMessage.Message = "Wrong password!"
		return responseMessage
	}
	responseMessage.Status = model.SUCCESS
	responseMessage.UserResp = userFromDB
	return responseMessage
}

// CreateUser add a new user in the DB
func CreateUser(form model.User) model.ResponseMessage {
	// Clean all white spaces
	model.RemoveSpaces(&form)

	if form.IsEmpty() {
		responseMessage.Status = model.FAILED
		responseMessage.Message = "All the fields are required"
		return responseMessage
	}

	if dao.Exists(form) {
		responseMessage.Status = model.FAILED
		responseMessage.Message = "The user already exists"
		return responseMessage
	}

	// Create the user
	dao.CreateUser(form)

	// Check if the new user has been successfully created
	if dao.Exists(form) {
		responseMessage.Status = model.SUCCESS
		responseMessage.Message = "The user has been sucessfully created"
	} else {
		responseMessage.Status = model.FAILED
		responseMessage.Message = "The registration failed"
	}
	return responseMessage
}
