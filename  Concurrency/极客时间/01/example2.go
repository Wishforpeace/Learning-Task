package main

import (
	"fmt"
	"sync"
)

// Mutex 嵌入字段使用
type Counter1 struct {
	sync.Mutex
	count uint64
}

func main() {
	var counter Counter1
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Lock()
				counter.count++
				counter.Unlock()
			}

		}()
	}
	wg.Wait()
	fmt.Println(counter.count)
}
