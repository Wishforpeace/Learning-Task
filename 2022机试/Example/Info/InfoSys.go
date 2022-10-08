package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net/http"
)

var DB *gorm.DB

type LoginRequest struct {
	ID string `json:"id"`
}
type Response struct {
	Code    int         `json:"code"`
	Message interface{} `json:"message"`
	Data    interface{} `json:"data"`
}

type Participator struct {
	gorm.Model
	Name      string
	Age       int
	StudentID string
	Gender    string
	College   string
}

func SendResponse(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code:    http.StatusOK,
		Message: message,
		Data:    data,
	})
}

func SendBadRequest(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(http.StatusBadRequest, Response{
		Code:    http.StatusBadRequest,
		Message: message,
		Data:    data,
	})
}

func SendError(c *gin.Context, message interface{}, data interface{}) {
	c.JSON(http.StatusInternalServerError, Response{
		Code:    http.StatusInternalServerError,
		Message: message,
		Data:    data,
	})
}

func Router() *gin.Engine {
	r := gin.New()
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API router.")
	})
	g1 := r.Group("/database")
	{
		g1.GET("/info", Info)
		g1.POST("/info", DatabaseInfo)
	}
	g2 := r.Group("/user")
	{
		g2.POST("/login", Login)
	}
	return r
}
func Info(c *gin.Context) {
	SendError(c, "非法入侵，警告警告⚠️", "❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️❗️")
}
func DatabaseInfo(c *gin.Context) {
	SendResponse(c, "成功获取数据库信息", gin.H{
		"name":     "exam_backend_2022",
		"address":  "rm-bp1pkcz9y269h30fm5o.mysql.rds.aliyuncs.com",
		"username": "fresh",
		"password": "Muxistudio517",
	})
}
func Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		SendBadRequest(c, "信息不足", nil)
		return
	}
	if err := CheckInfo(req.ID); err != nil {
		SendError(c, "信息验证失败", nil)
		return
	}
	SendResponse(c, "尊敬的贵宾您好，欢迎参加我们的晚会", nil)

}

func CheckInfo(id string) error {
	var par Participator
	return DB.Table("participators").Where("personal_id=?", id).Find(&par).Error
}

func main() {
	dsn := "fresh:Muxistudio517@tcp(rm-bp1pkcz9y269h30fm5o.mysql.rds.aliyuncs.com)/exam_backend_2022?charset=utf8mb4&parseTime=true&loc=Local"
	var err error
	DB, err = gorm.Open("mysql", dsn)
	if err != nil {
		fmt.Println("数据库连接失败")
		panic(err)
	}
	gin.SetMode(gin.ReleaseMode)
	r := Router()
	err = r.Run(":8080")
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
	log.Println("监听端口:", 8080, "请不要关闭终端")
	defer DB.Close()
}
