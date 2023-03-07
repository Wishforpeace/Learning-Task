package main

import (
	"fmt"
	"sync"
)

func foo() {
	var mu sync.Mutex
	defer mu.Unlock()
	fmt.Println("hello world")
}
func main() {
	foo()
}
