package dao

import (
	"database/sql"
	"sigmadoc/model"

	// MySQL open-source library
	_ "github.com/go-sql-driver/mysql"
)

// GetUser gets an user instance from the database
func GetUser(username string, password string) model.User {
	con, connectionError := sql.Open("mysql", "root:@/sigmadoc")
	defer con.Close()
	checkError(connectionError)

	row := con.QueryRow("SELECT * FROM user WHERE user_name = ? AND PASSWORD = ?", username, password)
	user := new(model.User)
	selectError := row.Scan(&user.Email, &user.Username, &user.Password)
	checkError(selectError)

	return *user
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
