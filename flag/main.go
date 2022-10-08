package main

import (
	"flag"
	"fmt"
)

var POTR = ""

func main() {
	path := flag.String("path", "Zhigui/Users", "文件的路径")
	POTR = *path
	fmt.Println(POTR)
}
