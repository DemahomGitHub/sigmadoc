package dao

import "database/sql"

// AbstractMapper is an interface that encapsulate the access to the DB
type AbstractMapper interface {
	Load(db *sql.DB, sqlQuery string) error
	Store(db *sql.DB, sqlQuery string) error
}

// Store stores an object in the database
func Store(m AbstractMapper, db *sql.DB, sqlQuery string) error {
	return m.Store(db, sqlQuery)
}

// Load loads an object from the database
func Load(m AbstractMapper, db *sql.DB, sqlQuery string) error {
	return m.Load(db, sqlQuery)
}
