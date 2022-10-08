package main

import (
	"fmt"
	"runtime"
	"time"
)

func showMsg() {
	for i := 0; i < 10; i++ {
		fmt.Printf("i:%v\n", i)
		if i >= 5 {
			runtime.Goexit()
		}
	}
}

func main() {
	go showMsg()
	time.Sleep(time.Second)
}
