package main

import "fmt"

func main() {
	ch := make(chan int)
	ch <- 10  //将10存入通道之中
	x := <-ch //从通道内接受值并赋给变量x
	//<-ch      //从ch中接受值，忽略结果
	fmt.Println(x)
	close(ch)
}
