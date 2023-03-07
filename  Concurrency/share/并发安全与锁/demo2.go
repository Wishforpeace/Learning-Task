package main

import (
	"fmt"
	"sync"
)

var x2 int64
var wg2 sync.WaitGroup
var lock sync.Mutex

func add2() {
	for i := 0; i < 50000; i++ {
		lock.Lock() // 加锁
		x2 = x2 + 1
		lock.Unlock() // 解锁
	}
	wg2.Done()
}
func main() {
	wg2.Add(2)
	go add2()
	go add2()
	wg2.Wait()
	fmt.Println(x2)
}
