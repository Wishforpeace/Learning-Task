package handle

import (
	"mygo/Mysql/Code/task3/model"
	"net/http"
)

var (
	username string
	password string
)

func Edit(w http.ResponseWriter, r *http.Request) {
	id := Get(w, r)
	_, ok := model.Find(id)
	if ok {
		username, password = Change(w, r)
		model.Update(id, username, password)
	}
}
