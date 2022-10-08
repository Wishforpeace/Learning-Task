package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 创建int类型通道，只能穿入int类型值
// 读写通道
var values1 = make(chan int)

// 只读通道
var values2 = make(<-chan int)

// 只写通道
var values3 = make(chan<- int)

// 可缓存的
var cha = make(chan int, 2)

func send() {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(10)
	fmt.Printf("send:%v\n", value)
	// time.Sleep(time.Second *5)
	values1 <- value
}
func main() {
	// 	从通道接收值
	defer close(values1)
	go send()
	fmt.Println("wait...")
	// 可读通道
	value := <-values1
	defer close(cha)
	cha <- 3
	cha <- 2
	fmt.Printf("receiver:%v\n", value)
	fmt.Println("end...")
	fmt.Println(len(cha))
	//fmt.Println("cha", <-cha)
	//fmt.Println("cha", <-cha)
	n := len(cha)
	for i := 0; i < n; i++ {
		fmt.Println("cha", <-cha)
	}

}
