package handle

import (
	"fmt"
	"Mygo/football/model"
	"net/http"
)

func AddDetails(w http.ResponseWriter, r *http.Request) {
	game := r.FormValue("GameName")
	details := r.FormValue("details")
	err := model.AddDetails(game, details)
	if err != nil {
		fmt.Println("信息修改失败")
	} else {
		fmt.Println("信息修改成功")
	}
}
