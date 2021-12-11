package model

import (
	"fmt"
	"mygo/Mysql/Code/UserSystem/utils"
)

type User struct {
	Uid   int
	Uname string
	Upwd  string
}

func CreatUser(name string, pwd string) error {
	sqlStr := "insert into live(userName,Pwd)values(?,?)"
	_, err := utils.DB.Exec(sqlStr, name, pwd) //将数据插入
	if err != nil {
		fmt.Println(err)
		return err
	} else {
		return nil
	}
}
func IsTrue(name string, pwd string) (string, string, error) {
	var username string
	var password string
	var err error
	sqlStr := "select userName,Pwd from live where userName =? and Pwd =?;"
	row, _ := utils.DB.Query(sqlStr, name, pwd)
	for row.Next() {
		err = row.Scan(&username, &password)
	}
	if err != nil {
		return "", "", err
	}
	return username, password, nil
}
