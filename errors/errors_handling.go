package errors

import (
	"log"
)

// CheckError :
func CheckError(err error, errorMessage string) error {
	if err != nil {
		log.Fatal(errorMessage, " Caused by ", err)
	}
	return err
}
