package dao

import (
	"database/sql"
	"log"
	mdl "sigmadoc/model"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	getUserFile    string = "/sql/get_user.sql"
	createUserFile string = "/sql/create_user.sql"
)

// GetUserByUsernameAndPassword gets a user instance from the database
func GetUserByUsernameAndPassword(form mdl.User) (mdl.User, error) {
	db := OpenDatabase()
	defer db.Close()

	sql := PrepareQuery(getUserFile)
	result, err := execGetUserQuery(db, sql, form.Username)
	return result, err
}
func execGetUserQuery(con *sql.DB, sqlQuery string, username string) (mdl.User, error) {
	row := con.QueryRow(sqlQuery, username)
	user := new(mdl.User)
	err := row.Scan(&user.Username, &user.Email, &user.Password)

	if err != nil {
		return *user, err
	}

	return *user, nil
}

// CreateUser creates a new user
func CreateUser(form mdl.User) {
	db := OpenDatabase()
	defer db.Close()

	sqlQuery := PrepareQuery(createUserFile)
	_, err := execCreateUserQuery(db, sqlQuery, form)

	if err != nil {
		log.Println("No user has been created")
		log.Println(err)
	}
}
func execCreateUserQuery(db *sql.DB, query string, form mdl.User) (sql.Result, error) {
	stmt, err := db.Prepare(query)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// Encrypt password before saving it
	hashPassword, err := bcrypt.GenerateFromPassword([]byte(form.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}
	res, err := stmt.Exec(form.Username, form.Email, string(hashPassword))

	if err != nil {
		return nil, err
	}

	return res, nil
}

// Exists check if a user exists
func Exists(user mdl.User) bool {
	res, err := GetUserByUsernameAndPassword(user)
	if err != nil {
		return false
	}
	usernameExists := strings.Compare(user.Username, res.Username) == 0
	emailExists := strings.Compare(user.Email, res.Email) == 0
	return usernameExists || emailExists
}
