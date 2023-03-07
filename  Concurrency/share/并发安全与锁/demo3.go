package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x3     int64
	wg3    sync.WaitGroup
	lock3  sync.Mutex
	rwlock sync.RWMutex
)

func write() {
	// lock.Lock()   // 加互斥锁
	rwlock.Lock() // 加写锁
	x3 = x3 + 1
	time.Sleep(10 * time.Millisecond) // 假设读操作耗时10毫秒
	rwlock.Unlock()                   // 解写锁
	// lock.Unlock()                     // 解互斥锁
	wg3.Done()
}

func read() {
	// lock.Lock()                  // 加互斥锁
	rwlock.RLock()               // 加读锁
	time.Sleep(time.Millisecond) // 假设读操作耗时1毫秒
	rwlock.RUnlock()             // 解读锁
	// lock.Unlock()                // 解互斥锁
	wg3.Done()
}

func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg3.Add(1)
		go write()
	}

	for i := 0; i < 1000; i++ {
		wg3.Add(1)
		go read()
	}

	wg3.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))
}
