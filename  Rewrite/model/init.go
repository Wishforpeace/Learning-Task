package model

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB
var err error

func Initdb() *gorm.DB {
	DB, err = gorm.Open("mysql", "root:root&1234@tcp(0.0.0.0:3306)/football?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println(err.Error())
	}
	return DB
}
