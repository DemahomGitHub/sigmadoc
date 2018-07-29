package dao

import (
	"database/sql"
	"os"
	errs "sigmadoc/errors"
	mdl "sigmadoc/model"

	// MySQL library for GO
	_ "github.com/go-sql-driver/mysql"
)

const getUserFileLocation string = "/sql/get_user.sql"

var errIns errs.ErrorInstance

// GetUserByUsernameAndPassword gets a user instance from the database
func GetUserByUsernameAndPassword(username string, password string) (mdl.User, error) {
	con := openConnection()
	defer con.Close()

	sql := prepareQuery()
	result, err := executeQuery(con, sql, username)
	return result, err
}

func openConnection() *sql.DB {
	con, err := sql.Open(MysqlDriver, MysqlDataSource)
	errs.CheckError(err, "Enable to open a connection with the database")
	return con
}

func prepareQuery() string {
	// Get working directory to find the sql file
	wd, _ := os.Getwd()
	location := wd + getUserFileLocation
	sql, err := GetQuery(location)
	errs.CheckError(err, "Enable to read the file containing the login query")
	return sql
}

func executeQuery(con *sql.DB, sqlQuery string, username string) (mdl.User, error) {
	row := con.QueryRow(sqlQuery, username)
	user := new(mdl.User)
	err := row.Scan(&user.Username, &user.Email, &user.Password)

	if err != nil {
		return *user, err
	}

	return *user, nil
}
