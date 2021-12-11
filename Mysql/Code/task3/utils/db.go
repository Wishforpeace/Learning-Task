package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func initDB() {
	DB, err := sql.Open("mysql", "root:root@(127.0.0.1)/usersystem")
	if err != nil {
		log.Println("Connection failed", err)
	} else {
		fmt.Println("Connection succeeded")
	}
	defer DB.Close()
}
