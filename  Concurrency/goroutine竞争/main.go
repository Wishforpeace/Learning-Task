package main

import (
	"fmt"
	"sync"
	"time"
)

var MyWg sync.WaitGroup

func Printer(str string) {
	for i := 0; i < len(str); i++ {
		//fmt.Println(str[i])
		fmt.Printf("%c", str[i])
		time.Sleep(time.Second)
	}

}

func Person1() {
	str := "hello"
	Printer(str)
	MyWg.Done()
}

func Person2() {

	str := "giao"
	Printer(str)
	MyWg.Done()
}
func main() {
	MyWg.Add(2)
	go Person1()
	go Person2()

}
