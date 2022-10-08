package main

import (
	"github.com/DeanThompson/ginpprof"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	ginpprof.Wrap(router)
	router.Run()
}
