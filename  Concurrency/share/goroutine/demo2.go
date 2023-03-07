package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello2(i int) {
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}

func main() {
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello2(i)
	}
	wg.Wait()
}
