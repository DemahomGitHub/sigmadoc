package service

import (
	"sigmadoc/dao"
	"sigmadoc/errors"
	mdl "sigmadoc/model"
	"strings"
)

var successResponse mdl.UserResponse
var errorResponse mdl.ErrorResponse

// GetUserByUsernameAndPassword gets a user from the DB
func GetUserByUsernameAndPassword(username string, password string) (mdl.UserResponse, mdl.ErrorResponse) {
	username = strings.TrimSpace(username)
	password = strings.TrimSpace(password)

	if len(username) == 0 {
		errorResponse.Status = errors.FAILED
		errorResponse.Info = "The User name is required"
		return successResponse, errorResponse
	}

	if len(password) == 0 {
		errorResponse.Status = errors.FAILED
		errorResponse.Info = "The Password is required"
		return successResponse, errorResponse
	}

	user, err := dao.GetUserByUsernameAndPassword(username, password)

	if err != nil {
		errorResponse.Status = errors.FAILED
		errorResponse.Info = "This user doesn't exist"
		return successResponse, errorResponse
	}

	if user.Equals(mdl.User{Username: username, Password: password}) {
		successResponse.Status = errors.SUCCESS
		successResponse.Data = user
	} else {
		errorResponse.Status = errors.FAILED
		errorResponse.Info = "Wrong user name or password"
	}
	return successResponse, errorResponse
}
