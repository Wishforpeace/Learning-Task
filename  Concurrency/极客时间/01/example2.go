package main

import (
	"fmt"
	"sync"
)

// Mutex 嵌入字段使用
type Counter1 struct {
	mu    sync.Mutex
	count int
}

func main() {
	var counter = Counter1{}
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.mu.Lock()
				counter.count++
				counter.mu.Unlock()
			}

		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}
