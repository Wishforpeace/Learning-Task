package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Step1:", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Step3", url)
	fmt.Println("Steps3", url)
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal("Step4", len(body))
	}
	fmt.Println("Step4:", len(body))
}

func main() {
	go responseSize("http://www.duoke360.com")
	go responseSize("http://baidu.com")
	go responseSize("http://jd.com")
	time.Sleep(10 * time.Second)
}
