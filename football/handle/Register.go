package handel

import (
	"fmt"
	"Mygo/football/model"
	"net/http"
)

func registrationsHandler(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.FormValue("username") == "" || r.FormValue("password") == "" {
		fmt.Fprintf(w, "Please enter a valid username and password.\r\n")
	} else {
		response, err := model.CreatUser(r.FormValue("username"), r.FormValue("password"), r.FormValue("identity"))
		if err != nil {
			fmt.Fprintf(w, err.Error())
		} else {
			fmt.Fprintln(w, response)
		}
	}
	UID = r.FormValue(identity)
	if err != nil {
		fmt.Println("err")
		return
	} else {
		fmt.Println(w, "success")
		switch UID {
		case "0":
			fmt.Println("最高管理员你好")
		case "1":
			fmt.Println("前台你好")
		case "2":
			fmt.Println("用户您好")
		default:
			break
		}
	}
}
