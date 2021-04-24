package db

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB
func InitDB() {
	d, err := sql.Open("sqlite3", "./test.db")
	if err != nil {
		panic(err)
	}
	db = d
}

func GetDB() *sql.DB {
	return db
}
