package model

import (
	"Mygo/football/utils"
	"log"
)

func AddDetails(game string, details string) error {
	results, err := utils.DB.Exec("UPDATE games SET info=? where game_name=?", details, game)
	if err != nil {
		log.Println("update data fail,err:", err)
	}
	log.Println(results.RowsAffected())
	return err
}
