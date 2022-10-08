package main

import (
	"log"
	"net/http"
)

func main() {
	request, err := http.NewRequest("GET", "https://account.ccnu.edu.cn/cas/login", nil)
	if err != nil {
		log.Println(err)
	}
	client := http.Client{
		//Timeout: TIMEOUT,
	}
	resp, err := client.Do(request)
	if err != nil {

		log.Println(err)

	}
	log.Println("resp", resp)
}
