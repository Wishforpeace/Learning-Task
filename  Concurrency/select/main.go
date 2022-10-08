package main

import "fmt"

func main() {
	ch := make(chan int)
	go func() {
		ch <- 3
	}()
	select {
	case num := <-ch:
		fmt.Println(num)
		fmt.Println("读取操作没问题")

	case ch <- 4:
		fmt.Println("写操作没问题")
	default:
		fmt.Println("default操作")
	}
}
