package handle

import (
	"fmt"
	"net/http"
	"test/model"
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
