package db

import "database/sql"

func Conn() *sql.DB {
	return db
}
