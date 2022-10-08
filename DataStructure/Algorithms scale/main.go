package main

import "fmt"

func sum(n int) int {
	total := 0
	for i := 0; i <= n; i++ {
		total = total + i
	}

	return total
}
func main() {
	fmt.Println(sum(100))
}
