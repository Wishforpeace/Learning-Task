package main

import (
	"fmt"
	"sync"
)

var sum int
var WGroup sync.WaitGroup

func Sum() {
	for i := 0; i < 10; i++ {
		fmt.Println(sum)
		sum++
	}
	WGroup.Done()
}

func main() {
	WGroup.Add(2)
	go Sum()
	go Sum()

	WGroup.Wait()
}
