package database

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(0.0.0.0:3306)/test")
	if err != nil {
		log.Print(err.Error())
	}
	return db
}

var DatabaseConnection *sql.DB = connect()
