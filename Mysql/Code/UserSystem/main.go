package main

import (
	"mygo/Mysql/Code/UserSystem/handle"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handle.Register)
	http.HandleFunc("/login", handle.Login)
	http.ListenAndServe(":10086", nil)
}
