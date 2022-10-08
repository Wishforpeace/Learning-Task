// package handle

// import (
// 	"fmt"
// 	"net/http"
// 	"test/model"
// )

// func RegistrationsHandler(w http.ResponseWriter, r *http.Request) {
// 	r.ParseForm()
// 	user := r.FormValue("username")
// 	password := r.FormValue("password")
// 	identity := r.FormValue("identity")
// 	if r.FormValue("username") == "" || r.FormValue("password") == "" {
// 		fmt.Fprintf(w, "Please enter a valid username and password.\r\n")
// 	} else {
// 		response, err := model.CreatUser(user, password, identity)
// 		if err != nil {
// 			fmt.Fprintf(w, err.Error())
// 			return

// 		}

// 		fmt.Fprintln(w, response)

// 	}
// 	UID := r.FormValue("identity")

// 	switch UID {
// 	case "0":
// 		fmt.Println("最高管理员你好")
// 	case "1":
// 		fmt.Println("前台你好")
// 	case "2":
// 		fmt.Println("用户您好")
// 	default:
// 		break
// 	}

// 	return
// }
