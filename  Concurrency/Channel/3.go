package main

import (
	"fmt"
	"time"
)

func hello(done chan bool) {
	fmt.Println("hello world goroutine")
	time.Sleep(4 * time.Second)
	fmt.Println("hello go routine awake and going to write to done")
	done <- true
}

func main() {
	done := make(chan bool)
	fmt.Println("Main going to call hello go routine")
	go hello(done)
	<-done
	fmt.Println("main received data")
}
