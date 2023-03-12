package main

import (
	"crawler/engine"
	"crawler/model"
	"crawler/scheduler"
	"crawler/zhenai/parser"
	"fmt"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	cfg = pflag.StringP("config", "c", "", "apiserver config file path.")
)

func main() {

	//engine.Run(engine.Request{
	//	Url:       "https://www.zhenai.com/zhenghun",
	//	ParseFunc: parser.ParseCityList,
	//})

	e := engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueueScheduler{},
		WorkerCount: 50,
	}
	e.Run(engine.Request{
		Url:       "https://www.zhenai.com/zhenghun",
		ParseFunc: parser.ParseCityList,
	})
}

func init() {
	var err error
	config := viper.New()

	config.AddConfigPath("conf/")  //设置读取的文件路径
	config.SetConfigName("config") //设置读取的文件名
	config.SetConfigType("yaml")   //设置文件的类型
	//尝试进行配置读取
	if err = config.ReadInConfig(); err != nil {
		panic(err)
	}

	//打印文件读取出来的内容:
	//fmt.Println(config.Get("db.addr"))
	//fmt.Println(config.Get("db.username"))
	//fmt.Println(config.Get("db.name"))
	//fmt.Println(config.Get("db.password"))
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.Get("db.username"),
		config.Get("db.password"),
		config.Get("db.addr"),
		config.Get("db.name"))

	model.DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	err = model.CreateTable()
	if err != nil {
		panic(err)
	}
}
