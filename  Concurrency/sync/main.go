package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func Hi2() {
	fmt.Println("执行Hi2")
	wg.Done()
}

func Hi1() {
	fmt.Println("执行Hi1")
	wg.Done()
}
func main() {
	start_time := time.Now()
	wg.Add(2)
	go Hi1()
	go Hi2()
	wg.Wait()
	end_time := time.Now()
	fmt.Println(end_time.Sub(start_time))
}
