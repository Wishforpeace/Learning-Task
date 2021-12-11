package main

import (
	"mygo/Mysql/Code/task3/handle"
	"net/http"
)

func main() {
	http.HandleFunc("/register", handle.Register)
	http.HandleFunc("/login", handle.Login)
	http.HandleFunc("/login/edit", handle.Edit)
	http.HandleFunc("/login/view", handle.View)
	http.HandleFunc("/login/delete", handle.Delete)
	err := http.ListenAndServe(":10086", nil)
	handle.Error(err)

}
