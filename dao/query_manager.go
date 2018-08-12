package dao

import (
	"database/sql"
	"io/ioutil"
	"os"
	errs "sigmadoc/errors"
	// MySQL library for GO
	_ "github.com/go-sql-driver/mysql"
)

const (
	// MysqlDriver : name of the driver
	MysqlDriver string = "mysql"
	// MysqlDataSource : information about the database
	MysqlDataSource string = "root:FHcc3Hjr@/sigmadoc"
)

// OpenDatabase :
func OpenDatabase() *sql.DB {
	db, err := sql.Open(MysqlDriver, MysqlDataSource)
	errs.CheckError(err, "Enable to open a connection with the database")
	return db
}

// PrepareQuery :
func PrepareQuery(fileLocation string) string {
	// Get working directory to find the sql file
	wd, _ := os.Getwd()
	sql, err := GetQuery(wd + fileLocation)
	errs.CheckError(err, "Enable to read the file containing the query")
	return sql
}

// GetQuery reads a SQL file and returns
// the content of the file in a string format
// or an error when something went wrong
func GetQuery(location string) (string, error) {
	sql, err := ioutil.ReadFile(location)
	if err != nil {
		return "", err
	}
	return string(sql), nil
}
