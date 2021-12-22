package model

import (
	"Mygo/football/utils"
	"log"
)

func Who(name string, pwd string) int {
	var idnum int
	sqlStr := "select position from users where user_name = ?"
	rows, err := utils.DB.Query(sqlStr, name, pwd)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&idnum)
		if err != nil {
			log.Fatal(err)
		}

	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
	return idnum
}
