package handle

import (
	"fmt"
	"mygo/Mysql/Code/UserSystem/model"
	"net/http"
)

func Login(w http.ResponseWriter, r *http.Request) {
	Uname := r.FormValue("username")
	Upwd := r.FormValue("password")
	name, _, err := model.IsTrue(Uname, Upwd)
	if err != nil {
		fmt.Println(w, "fail")
		fmt.Println(err)
	} else {
		fmt.Fprintln(w, "success")
		fmt.Println(name)
	}

}
