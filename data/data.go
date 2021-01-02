package data

import (
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

// Db Global Reference To Database
var Db *sql.DB

func init() {
	log.Println("Initializing DB")
	var err error
	Db, err = sql.Open("mysql", "root:password@tcp(127.0.0.1:3306)/social_connect")
	if err != nil {
		panic(err.Error())
	}
	return
}