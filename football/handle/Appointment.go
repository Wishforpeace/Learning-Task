package handel

import (
	"Mygo/football/model"
	"fmt"
	"net/http"
)

func Appointment(w http.ResponseWriter, r *http.Request) {
	game := r.FormValue("GameName")
	err := model.Reserve(game)
	if err != nil {
		fmt.Println(w, err)
		fmt.Println("预约失败")
	} else {
		fmt.Println("预约成功")
	}
}
