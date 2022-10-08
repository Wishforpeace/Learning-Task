package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/student/0509/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// @title Swagger Example API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService https://www.topgoer.com

// @contact.name www.topgoer.com
// @contact.url https://www.topgoer.com
// @contact.email me@razeen.me

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host 127.0.0.1:8080
// @BasePath /api/v1

func main() {

	r := gin.Default()

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/hello", HandleHello)
		// v1.POST("/login", HandleLogin)
		// v1Auth := r.Use(HandleAuth)
		// {
		//     v1Auth.POST("/upload", HandleUpload)
		//     v1Auth.GET("/list", HandleList)
		// }
	}

	r.Run(":8080")
}
