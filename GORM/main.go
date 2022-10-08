package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strconv"
)

type User struct {
	gorm.Model
	Email      string    `json:"email" gorm:"varchar(100);not null;unique"`
	NickName   string    `json:"nick_name"gorm:"varchar(20);not null;unique"`
	Password   string    `json:"password" gorm:"required;not nnull"`
	Gender     string    `json:"gender" `
	Degree     string    `json:"degree"  gorm:"varchar(10);not null"`
	Publisheds []Task    `gorm:"foreignKey:Pubilsher;reference:Email"`
	Accpeteds  []Task    `gorm:"foreignKey:Accepter;reference:Email"`
	Payment    Payment   `gorm:"foreignKey:Email;reference:Email"`
	Image      UserImage `gorm:"foreignKey:owner;reference:email"`
	Earning    string    `gorm:"varchar(10)"`
}
type Task struct {
	gorm.Model
	Content   string     `json:"content" gorm:"varchar(255)"`
	Publisher string     `json:"publisher" gorm:"unique;<-:create"`
	Accepter  string     `json:"accepter"gorm:"unique;<-:create"`
	Award     string     `json:"award"`
	Status    string     `json:"status" gorm:"column:status;varchar(10)"`
	Image     TaskImages `gorm:"foreignKey:TaskID;reference:ID"`
}
type UserImage struct {
	gorm.Model
	Owner  string `gorm:"<-:create;type:varchar(30)"`
	Avatar string
	Sha    string
	Path   string
}
type TaskImages struct {
	gorm.Model
	TaskID uint
	Avatar string
	Sha    string
	Path   string
}
type Payment struct {
	gorm.Model
	Email  string `gorm:"<-:create;type:varchar(100)"`
	Avatar string
	Sha    string
	Path   string
}

func main() {
	dsn := "root:root&1234@tcp(127.0.0.1:3306)/Zhigui?charset=utf8mb4&parseTime=True&loc=Local"
	DB, err := gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	// DB.AutoMigrate(&Zone{}, &Prefecture{})
	// var zone Zone

	// // pre1 := Prefecture{
	// // 	Name: "篮球",
	// // }
	// // pre2 := Prefecture{
	// // 	Name: "足球",
	// // }
	// zone = Zone{
	// 	Name: "体育",
	//  }
	// // DB.Create(&zone)
	// // DB.Debug().Model(&Zone{}).Save(&zone)
	// var pre []Prefecture

	//user.Email = "1903180340@qq.com"
	//var Image UserImage
	//DB.Debug().Preload("UserImage", "owner=?", user.Email).First(&Image)
	//fmt.Printf("Image:%v", Image)
	// DB.Debug().Preload("Prefectures").First(&z)
	// fmt.Println(z.Name)

	offset := 0
	limit := 5
	prefecture := 1
	item := make([]*Task, 0)
	d := DB.Table("tasks").
		Where(" prefecture_id = ?", prefecture).
		Offset(offset).Limit(limit).
		Order("created_at").Find(&item)
	if d.Error != nil {
		log.Println(d.Error)
	}
	fmt.Println(item)
	id := "22222"
	Id, _ := strconv.ParseUint(id, 10, 32)
	ID := uint(Id)
	fmt.Printf("%T", ID)
}
