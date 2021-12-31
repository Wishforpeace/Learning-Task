package handle

import (
	"fmt"
	"net/http"
	"strings"
	"test/model"
)

func TestResourceHandler(w http.ResponseWriter, req *http.Request) {

	authToken := strings.Split(req.Header.Get("Authorization"), "Bearer ")[1]

	userDetails, err := model.ValidateToken(authToken)

	if err != nil {

		fmt.Fprintf(w, err.Error())

	} else {

		username := fmt.Sprint(userDetails["username"])

		fmt.Fprintf(w, "Welcome, "+username+"\r\n")
	}

}
