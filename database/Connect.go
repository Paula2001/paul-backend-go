package database

import (
	"awesomeProject/Env"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"os"
)

func connect() *sql.DB {
	Env.LoadEnv() // Todo : need to create bootstrap for the project
	var name string = os.Getenv("MYSQL_USER_NAME")
	var host string = os.Getenv("MYSQL_HOST")
	var port string = os.Getenv("MYSQL_PORT")
	var database string = os.Getenv("MYSQL_DATABASE_NAME")
	var password string = os.Getenv("MYSQL_PASS")

	println(os.Getenv("MYSQL_USER_NAME"), "------", "asd")
	db, err := sql.Open("mysql", name+":"+password+"@tcp("+host+":"+port+")/"+database)

	if err != nil {
		panic(err.Error())
	}
	return db
}

var Connection *sql.DB = connect() // Todo : this need to be singleton or should close after we finish
