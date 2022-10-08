package main

import (
	"fmt"

	"sync"
)

func main() {
	var counter Counter2
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 1000; j++ {
				counter.Incr()
			}
		}()
	}
	wg.Wait()
	fmt.Println(counter.Count())
}

// 线性安全的计数器形式
type Counter2 struct {
	CounterType int
	Name        string
	mu          sync.Mutex
	count       uint64
}

// 加1的方法，内部使用互斥锁保护
func (counter *Counter2) Incr() {
	counter.mu.Lock()
	counter.count++
	counter.mu.Unlock()
}

//得到计数器的值，也需要互斥锁的保护
func (counter *Counter2) Count() uint64 {
	counter.mu.Lock()
	defer counter.mu.Unlock()
	return counter.count
}
