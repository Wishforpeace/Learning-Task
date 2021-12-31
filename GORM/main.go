package main

import (
	"errors"
	"fmt"
	"os/user"
	"time"

	"Error"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var(
	DB *gorm.DB
)
func main() {

	dsn := "root:root&1234@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	} else {
		fmt.Print("Connected")
	}
	type User struct {
		ID       int
		Name     string
		Age      uint8
		Birthday time.Time
	}
	user := User{Name: "Jinzhu", Age: 18, Birthday: time.Now()}
	result := db.Create(&user)
	user.ID
	result.Error
	result.RowsAffected
	//用指定字段创建记录
	DB.Select("Name", "Age", "CreateAt").Create(&user)
	//INSERT INTO `users` (`name`,`age`,`created_at`) VALUES ("jinzhu", 18, "2020-07-04 11:05:21.775")
	//创建一个记录且一同忽略传递给略去的字段值
	DB.Omit("Name", "Age", "CreatedAt").Create(&user)
	//批量插入
	var users = []User{{Name: "jinzhu1"}, {Name: "jinzhu2"}, {Name: "jinzhu3"}}
	DB.Create(&users)
	for _, user := range users {
		user.ID
	}
	//使用CreateInBatches分批创建时，你可以指定没批的数量，例如：
	var users =[]User{{Name:"jinzhu"},{Name:"jinzhu_1000"}}
	DB.CreateInBaches(users,100)
	DB.err = gorm.Open(sqlite.Open("gorm.db"),&gorm.Config{
		CreateBatchSize:1000,
	})
	DB = DB.Session(&gorm.Session{CreateBatchSize:1000})

	users = [5000]User{{Name:"jinzhu",Pets:[]Pet{pet1,pet2,pet3}}
	DB.Create(&users)




}

	//创建hook
	func (u *User) BeforeCreate(tx *gorm.DB) (err error) {
		u.UUID = uuid.New()
	  
		if !u.IsValid() {
		  err = errors.New("can't save invalid data")
		}
		return
	  }
	  
	  func (u *User) AfterCreate(tx *gorm.DB) (err error) {
		if u.ID == 1 {
		  tx.Model(u).Update("role", "admin")
		}
		return
	  }

	  func (u *User) AfterCreate(tx *gorm.DB) (err error) {
		if !u.IsValid() {
		  return errors.New("rollback invalid user")
		}
		return nil
	  }