package router

import (
	"github.com/gin-gonic/gin"
	"myproject/handler"
)

func Router(r *gin.Engine) {
	r.POST("/user", handler.User)
	r.POST("/login", handler.Login)
	//球赛
	v2 := r.Group("/football")

}
