package main

import (
	"fmt"
)

func main() {
	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("world")

	// 主协程
	for i := 0; i < 2; i++ {
		//runtime.Gosched()
		fmt.Println("hello")
	}
}
