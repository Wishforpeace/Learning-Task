package main

import (
	"net/http"
	"test/handle"
)

func main() {
	http.HandleFunc("/register", handle.RegistrationsHandler)
	http.HandleFunc("/auth", handle.AuthenticationsHandler)
	http.HandleFunc("/test", handle.TestResourceHandler)
	// http.HandleFunc("/gamerange", handle.RangeGames)
	http.HandleFunc("/add_palyes", handle.EditPlayers)
	http.ListenAndServe(":8081", nil)
}
