package main

import (
	"Mygo/football/handle"
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/registeration", handle.registrationsHandler)
	http.HandleFunc("/authentications", handle.authenticationsHandler)
	http.HandleFunc("/test", handle.testResourceHandler)
	http.HandleFunc("/gamerange", handle.RangeGames)
	http.HandleFunc("/AddPalyes", handle.EditPlayers)
	http.ListenAndServe(":8081", nil)
}
