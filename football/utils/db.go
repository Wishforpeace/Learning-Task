package utils

import (
	"database/sql"
	"fmt"
)

var DB *sql.DB
var err error

func InitDB() {
	DB, err = sql.Open("mysql", "root:root&1234@tcp(192.168.50.166:3306)/football")
	if err != nil {
		fmt.Println("connect to mysql failed,", err)
		return
	}
	fmt.Println("connect to mysql success")

	err = DB.Ping()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("ping success")
}
