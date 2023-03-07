package main

import "fmt"

func sendData(sendch chan<- int) {
	sendch <- 10
}
func main() {
	// 单向通道
	sendch := make(chan int)
	go sendData(sendch)
	fmt.Println(<-sendch)
}
