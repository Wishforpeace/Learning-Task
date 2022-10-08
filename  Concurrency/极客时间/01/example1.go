package main

import (
	"fmt"
	"sync"
)

func main() {
	// 互斥锁保护计时器
	var mu sync.Mutex
	// 计数器的值
	var count = 0

	// 辅助变量，用来确定所有goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)

	// 启动10个goroutine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			for j := 0; j < 1000; j++ {
				count++
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(count)
}
