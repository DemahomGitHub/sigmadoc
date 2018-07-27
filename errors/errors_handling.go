package errors

import (
	"fmt"
	"log"
)

// ErrorInstance represents an instance or error that occurs at a certain time
type ErrorInstance struct {
	ErrMessage string `json:"ErrorMessage"`
	ErrCause   error  `json:"ErrorCause"`
}

const (
	// SUCCESS is a response type
	SUCCESS int = 200
	// FAILED is a response type
	FAILED int = 400
)

// CheckError :
func CheckError(err error, errorMessage string) error {
	if err != nil {
		log.Fatal(errorMessage, " Caused by ", err)
	}
	return err
}

// HasErrors :
func (e ErrorInstance) HasErrors() bool {
	return e.ErrCause != nil
}

func (e *ErrorInstance) Error() string {
	return fmt.Sprintf("[%s] %v", e.ErrMessage, e.ErrCause)
}
