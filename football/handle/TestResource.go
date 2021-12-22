package handle

import (
	"Mygo/football/model"
	"fmt"
	"net/http"
	"strings"
)

func testResourceHandler(w http.ResponseWriter, req *http.Request) {

	authToken := strings.Split(req.Header.Get("Authorization"), "Bearer ")[1]

	userDetails, err := model.val(authToken)

	if err != nil {

		fmt.Fprintf(w, err.Error())

	} else {

		username := fmt.Sprint(userDetails["username"])

		fmt.Fprintf(w, "Welcome, "+username+"\r\n")
	}

}
