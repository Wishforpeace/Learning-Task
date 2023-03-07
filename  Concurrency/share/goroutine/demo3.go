package main

import (
	"fmt"
	"time"
)

func main() {

	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second * 2)
		}
	}()
	i := 0

	fmt.Printf("main goroutine: i = %d\n", i)
	time.Sleep(time.Second)

}
