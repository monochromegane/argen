package tests

import "database/sql"

var db *sql.DB

func Use(DB *sql.DB) {
	db = DB
}
