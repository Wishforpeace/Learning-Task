package main

import (
	"fmt"
	"runtime"
)

func show(c string) {
	for i := 0; i < 2; i++ {
		fmt.Printf("msg:%v", c)
	}
}

func main() {
	go show("java") // 子协程
	// 主协程
	for i := 0; i < 2; i++ {
		runtime.Gosched() // 我有权利执行任务了，让给其他子协程来执行
		fmt.Printf("\"golang\":%v\n", "golang")
	}
	fmt.Println("end...")
}
