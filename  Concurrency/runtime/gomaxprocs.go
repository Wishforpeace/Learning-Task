package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 0; i < 10; i++ {
		fmt.Println("a:", i)
	}
}

func b() {
	for i := 0; i < 10; i++ {
		fmt.Println("b:", b)
	}
}

func main() {
	fmt.Printf("runtime.NumCPU():%v\n", runtime.NumCPU())
	go a()
	go b()
	time.Sleep(time.Second)
}
