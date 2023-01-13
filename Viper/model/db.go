package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"test/log"
)

type Database struct {
	Self *gorm.DB
}

var DB *Database

func openDB(username, password, addr, name string) *gorm.DB {
	config := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=%t&loc=%s",
		username,
		password,
		addr,
		name,
		true,
		"Local")
	fmt.Println("config", config)

	db, err := gorm.Open("mysql", config)

	if err != nil {
		log.Error("Open database failed",
			zap.String("reason", err.Error()),
			zap.String("detail", fmt.Sprintf("Database connection failed.Database name: %s", name)))
	}

	setupDB(db)

	return db
}

func setupDB(db *gorm.DB) {
	db.LogMode(viper.GetBool("gormlog"))
	db.DB().SetMaxOpenConns(20000) // 用于设置最大连接数，默认为0，表示不限制，设置最大连接数，可以避免并发太高导致连接mysql出现too many connections的错误
	db.DB().SetConnMaxIdleTime(0)

	// 开发时打开SQL log
	if viper.GetString("runmode") == "debug" {
		db.LogMode(true)
	}
}

func InitSelfDB() *gorm.DB {
	return openDB(viper.GetString("db.username"),
		viper.GetString("db.password"),
		viper.GetString("db.addr"),
		viper.GetString("db.name"))
}

func GetSelfDB() *gorm.DB {
	return InitSelfDB()
}

func (db *Database) Init() {
	DB = &Database{
		Self: GetSelfDB(),
	}

}

func (db *Database) Close() {
	DB.Self.Close()
}
