package main

import (
	"fmt"
	"sync"
)

func hello(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	fmt.Println("hello", i)
	mu.Unlock()

}

var mu sync.Mutex

func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	for i := 0; i < 5; i++ {
		mu.Lock()
		go hello(&wg, i)

	}
	wg.Wait()
}
