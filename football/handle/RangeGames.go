package handle

import (
	"fmt"
	"net/http"
)

func Rangegames(w http.ResponseWriter, r *http.Request) {
	players := r.FormValue("palyers")
	time := r.FormValue("time")
	heat := r.FormValue("heat")
	_, err := model.RangePlayers(players)
}
