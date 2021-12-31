package handle

import (
	"fmt"
	"net/http"
	"test/model"
)

func SetGame(w http.ResponseWriter, r *http.Request) {
	NewGame := r.FormValue("newgame")
	Place := r.FormValue("place")
	GameData := r.FormValue("data")
	appointment := 0
	TeamA := r.FormValue("teamA")
	TeamB := r.FormValue("teamB")
	_, err := model.SetGame(NewGame, Place, GameData, appointment, TeamA, TeamB)
	if err != nil {
		fmt.Println(w, "设置失败")
	} else {
		fmt.Println("设置成功")
	}
}
