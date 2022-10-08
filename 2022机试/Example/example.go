package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Mysql struct {
	username string
	password string
	address  string
	name     string
}

type Participator struct {
	gorm.Model
	Name         string
	Age          int
	PersonalID   string
	Gender       string
	Organization string
}

var MysqlInfo Mysql

func openDB() (*gorm.DB, error) {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		MysqlInfo.username,
		MysqlInfo.password,
		MysqlInfo.address,
		MysqlInfo.name,
		true,
		"Local")
	fmt.Println(config)
	db, err := gorm.Open(mysql.Open(config), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return db, err
}
func main() {
	db, err := openDB()
	fmt.Println(err)
	db.AutoMigrate(&Participator{})
	if err = db.Model(&Participator{}).Create(&Participator{
		Name:         "test",
		Age:          1,
		PersonalID:   "2021216666",
		Gender:       "男",
		Organization: "后端组",
	}).Error; err != nil {
		panic(err)
		return
	} else {
		fmt.Println("创建成功")
	}
	var par Participator
	if err = db.Model(&Participator{}).Where("name = ?", "test").Find(&par).Error; err != nil {
		panic(err)
	}
	fmt.Println(par)
}

func init() {
	MysqlInfo = Mysql{
		name:     "exam_backend_2022",
		address:  "rm-bp1pkcz9y269h30fm5o.mysql.rds.aliyuncs.com",
		username: "fresh",
		password: "Muxistudio517",
	}
}
