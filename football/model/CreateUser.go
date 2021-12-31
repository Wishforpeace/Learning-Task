package model

import (
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
	"strconv"
	"test/utils"
)

func CreateUser(name string, pwd string, position string) (string, error) {
	sqlStr := "insert into users(user_name,user_password,position)values(?,?,?)"
	identity, _ := strconv.Atoi(position)
	stmt, err := utils.DB.Prepare(sqlStr)
	if err != nil {
		return "", err
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	_, err = stmt.Exec(name, hashedPassword, identity)
	if err != nil {
		return "", err
	}
	return "success\r\n", nil
}
