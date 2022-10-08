package main

import (
	"fmt"
	"time"
)

func showMsg(msg string) {
	for i := 0; i < 5; i++ {
		fmt.Printf("msg:%v\n", msg)
		//time.Sleep(time.Millisecond * 100)
	}
}
func main() {
	go showMsg("Java ") //go 启动了一个协程来执行1
	time.Sleep(time.Millisecond * 2000)
	go showMsg("golang") // 2
	time.Sleep(time.Millisecond * 2000)
	fmt.Println("main end ...") //3 主函数退出，程序结束
}
