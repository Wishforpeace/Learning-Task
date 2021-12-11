//获取Cookie
package handle

import (
	"fmt"
	"net/http"
)

func Get(w http.ResponseWriter, r *http.Request) string {
	cookie, err := r.Cookie("cookie")
	if err != nil {
		fmt.Println("获取cookie失败！", err)
	}
	return cookie.Value
}
