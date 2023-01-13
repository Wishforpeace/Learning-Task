package main

import "fmt"

func main() {
	var a *int
	*a += 1
	DoOperation(1, decrease)
}

func increase(a, b int) int {
	return a + b
}

func DoOperation(y int, f func(int, int)) {
	f(y, 1)
}

func decrease(a, b int) {
	fmt.Println("decrease result is:", a-b)
}
