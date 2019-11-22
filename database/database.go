package database

import "database/sql"

// Database interface
type Database interface {
	DB() *sql.DB
	DBTest() *sql.DB
}

// Impl database implementation
type Impl struct{}

// DBTest interface
type DBTest interface {
	Clean(tables ...string)
}

// DBTestImpl database test implementation struct
type DBTestImpl struct {
	Conn *sql.DB
}
