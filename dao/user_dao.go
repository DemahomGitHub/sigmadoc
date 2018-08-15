package dao

import (
	"database/sql"
	"log"
	"sigmadoc/model"
	"strings"

	"golang.org/x/crypto/bcrypt"
)

const (
	getUserFile    string = "/sql/get_user.sql"
	createUserFile string = "/sql/create_user.sql"
)

// UserImpl :
type UserImpl model.User

// GetUserByUsernameAndPassword gets a user instance from the database
func GetUserByUsernameAndPassword(form model.User) (model.User, error) {
	db := OpenDatabase()
	defer db.Close()

	sql := PrepareQuery(getUserFile)
	usr := UserImpl(form)
	err := Load(&usr, db, sql)
	return model.User(usr), err
}

// Load is the implementation of a method in AbstractMapper interface
func (user *UserImpl) Load(db *sql.DB, sqlQuery string) error {
	row := db.QueryRow(sqlQuery, user.Username)
	err := row.Scan(&user.Username, &user.Email, &user.Password)
	if err != nil {
		return err
	}
	return nil
}

// CreateUser creates a new user
func CreateUser(form model.User) {
	db := OpenDatabase()
	defer db.Close()

	query := PrepareQuery(createUserFile)
	user := UserImpl(form)
	err := Store(&user, db, query)

	if err != nil {
		log.Println("No user has been created")
		log.Println(err)
	}
}

// Store is the implementation of the method in AbstractMapper interface
func (user *UserImpl) Store(db *sql.DB, query string) error {
	stmt, err := db.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(user.Username, user.Email, string(hashPassword))
	return err
}

// Exists check if a user exists
func Exists(user model.User) bool {
	res, err := GetUserByUsernameAndPassword(user)
	if err != nil {
		return false
	}
	usernameExists := strings.Compare(user.Username, res.Username) == 0
	emailExists := strings.Compare(user.Email, res.Email) == 0
	return usernameExists || emailExists
}
