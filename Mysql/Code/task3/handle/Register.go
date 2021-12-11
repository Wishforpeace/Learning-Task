package handle

import (
	"fmt"
	"mygo/Mysql/Code/task3/model"
	"net/http"
)

func Register(w http.ResponseWriter, r *http.Request) {
	Uname := r.FormValue("username")
	Upwd := r.FormValue("password")
	err := model.CreatUser(Uname, Upwd)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(w, "success")
	}
}
