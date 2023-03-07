package main

import "fmt"

func main() {
	ch := make(chan int, 10) // 创建一个容量为1的有缓冲区通道
	ch <- 10
	fmt.Println("发送成功")
	fmt.Println(len(ch))
	fmt.Println(cap(ch))
}
