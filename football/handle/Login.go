package handle

import (
	"fmt"
	"net/http"
	"test/model"
)

func Login(w http.ResponseWriter, r *http.Request) error {
	Uname := r.FormValue("username")
	Upwd := r.FormValue("password")
	_, err := model.IsTrue(Uname, Upwd)
	identity := model.Who(Uname, Upwd)
	if err != nil {
		fmt.Println(w, "fail")
		fmt.Println(err)
	} else {
		switch identity {
		case 0:
			fmt.Println("登录成功,您的身份是（主管）")
		case 1:
			fmt.Println("登录成功，您的身份为(前台人员)")
		case 2:
			fmt.Println("登录成功，您的身份为(普通用户)")
		default:
			break
		}
	}
	return err
}
