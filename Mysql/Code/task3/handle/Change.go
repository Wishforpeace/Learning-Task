package handle

import (
	"fmt"
	"net/http"
)

func Change(w http.ResponseWriter, r *http.Request) (string, string) {
	fmt.Println(w, "请输入更改信息")
	username := r.FormValue("userame")
	password := r.FormValue(("password"))
	return username, password
}
