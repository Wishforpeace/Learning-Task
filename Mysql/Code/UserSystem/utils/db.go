package utils

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var (
	DB  *sql.DB
	err error
)

func initDB() {
	DB, err := sql.Open("mysql", "root:root@/usersystem")
	if err != nil {
		log.Println(err)
	}
	log.Println(DB)

}
