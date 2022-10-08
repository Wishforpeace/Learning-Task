package main

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type TaskRequest struct {
	PrefectureID string `json:"prefecture_id"`
	Content      string `json:"content"`
	Award        string `json:"award"`
}

func main() {
	r := gin.Default()
	Router(r)
	r.Run(":4000")
}

func Router(r *gin.Engine) {
	g1 := r.Group("/api/v1/user")
	{
		g1.POST("/demo", Upload)
	}
}

func Upload(c *gin.Context) {
	task, true := c.GetPostForm("task")
	if !true {
		c.JSON(http.StatusNotFound, gin.H{
			"error": true,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "芜湖",
		})
	}
	U := []byte(task)
	m := TaskRequest{}
	err := json.Unmarshal(U, &m)
	if err != nil {
		log.Println("寄了", err)
	} else {
		log.Println("牛逼", m)
	}
}
