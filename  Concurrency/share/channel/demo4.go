package main

import "fmt"

func receiver(c chan int) {
	ret := <-c
	fmt.Println("接收成功", ret)
}
func main() {
	ch := make(chan int)
	go receiver(ch) // 启用goroutine从通道接收值
	ch <- 10
	fmt.Println("发送成功")
}
