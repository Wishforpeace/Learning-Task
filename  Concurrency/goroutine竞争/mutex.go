package main

import (
	"fmt"
	"sync"
	"time"
)

var sum1 int

//var WGroup1 sync.WaitGroup
var rwmutex sync.RWMutex

func Write() {
	rwmutex.Lock()
	sum1++
	//WGroup1.Done()
	fmt.Printf("写入数据正常，加完后值是：%d", sum1)
	rwmutex.Unlock()
}
func Read() {
	rwmutex.Lock()
	fmt.Println(sum1)
	//WGroup.Done()
	fmt.Printf("读数据正常，读到的值是%d", sum1)
	rwmutex.Unlock()
}

func main() {
	//WGroup1.Add(2)
	for i := 0; i < 10; i++ {
		go Write()
	}

	for j := 0; j < 10; j++ {
		go Read()
	}
	time.Sleep(20 * time.Second)
	//WGroup1.Wait()
}
