package main

import (
	"fmt"
	"sync"
)

var once sync.Once

func Hallen() {
	fmt.Println("hello world")
}
func main() {
	for i := 0; i < 10; i++ {
		once
		Hallen()
	}
}
