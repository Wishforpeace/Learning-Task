package model

import (
	"log"
	"test/utils"
)

func Reserve(game string) error {
	sqlStr := "update games set appointment = appointment+1 where game_name=?"
	_, err := utils.DB.Exec(sqlStr, game)
	if err != nil {
		log.Println("Reserve failed")
		log.Fatal(err)
	} else {
		err = nil
	}
	return err
}
