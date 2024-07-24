package db

import "database/sql"

type Connection struct {
	Server   string
	User     string
	Pass     string
	DataBase string
	Driver   string
	conn     *sql.DB
}
