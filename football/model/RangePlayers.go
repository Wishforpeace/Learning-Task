// package model

// import (
// 	"log"
// 	"test/utils"
// )

// func RangePlayers(name string) (string, error) {
// 	var team string
// 	sqlStr := "select team_name from players where name = ?"
// 	rows, err := utils.DB.Exec(sqlStr, team)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer rows.Close()
// 	for rows.Next() {
// 		err := rows.Scan(&team)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
// 	}
// 	err = rows.Err()
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	sqlStr = "select game_id,game_name from games where "
// }
