package dao

import (
	"io/ioutil"
)

const (
	// MysqlDriver : name of the driver
	MysqlDriver string = "mysql"
	// MysqlDataSource : information about the database
	MysqlDataSource string = "root:@/sigmadoc"
)

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
