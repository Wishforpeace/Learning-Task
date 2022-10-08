package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB
var err error

type User struct {
	ID        string `gorm:"column:id;AUTO_INCREMENT "`
	StudentID string `gorm:"column:student_id"`
	Name      string `gorm:"column:name"`
}

func InitDB() *gorm.DB {
	DB, err = gorm.Open("mysql", "root:root&1234@tpc(127.0.0.1:3306)/library?charset=utf-8mb4&parseTime=True&loc=local")
	if err != nil {
		fmt.Printf("%s", err)

	}
	DB.AutoMigrate(&User{})
	return DB
}
