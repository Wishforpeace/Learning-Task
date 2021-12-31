package model

import (
	"log"
	"test/utils"
)

func SetGame(NewGame string, Place string, data string, appointment int, teamA string, teamB string) (int64, error) {
	sqlStr := "INSERT INTO games(game_name,game_date,place,appointment,TEAMA,TEAMB)VALUES (?,?,?,?,?,?)"
	result, err := utils.DB.Exec(sqlStr, NewGame, Place, data, appointment, teamA, teamB)
	if err != nil {
		log.Println("insert failed", err)
	}
	lastId, err := result.LastInsertId()
	if err != nil {
		log.Fatal(err)
	}
	rowCnt, err := result.RowsAffected()
	if err != nil {
		log.Fatal(err)

	}
	log.Printf("ID = %d, affected = %d\n", lastId, rowCnt)
	return lastId, err
}
