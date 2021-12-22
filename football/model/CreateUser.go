package model

import (
	"Mygo/football/utils"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"golang.org/x/crypto/bcrypt"
)

func CreateUser(name string, pwd string, position int) (string, error) {
	sqlStr := "insert into users(user_name,user_password,position)values(?,?,?)"
	stmt, err := utils.DB.Exec(sqlStr, name, pwd, position)
	if err != nil {
		return "", err
	}
	defer stmt.Close()
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(pwd), 14)
	_, err = stmt.Exec(name, hashedPassword)
	if err != nil {
		return "", err
	}
	return "success\r\n", nil
}
