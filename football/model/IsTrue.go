package model

import (
	"log"
	"test/utils"
)

func IsTrue(Uname string, Upwd string) (string, error) {
	var err error
	sqlStr := "selcet user_name,user_password from live where user_name = ? and user_password =?"
	rows, err := utils.DB.Query(sqlStr, Uname, Upwd)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&Uname, &Upwd)
		if err != nil {
			log.Fatal(err)
			Uname = ""
			err = nil
		}
	}
	return Uname, err
}
