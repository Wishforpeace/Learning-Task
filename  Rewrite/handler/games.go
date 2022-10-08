package handler

import (
	"github.com/gin-gonic/gin"
	"myproject/model"
)

func ViewGames(c *gin.Context) {
	var games []model.Games
	num := 10
	model.DB.Find(&games)

	for i := len(games) - 1; i >= 0; i-- {
		if num > 0 {
			num--
		} else {
			break
		}

	}
}
