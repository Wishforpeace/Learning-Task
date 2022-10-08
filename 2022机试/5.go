package main

import (
	"fmt"
	"math"
)

func Reverse(num int) int {
	var binary []int
	for num != 0 {
		binary = append(binary, num%2)
		num = num / 2
	}
	reverse_num := 0
	for i, m := range binary {
		//fmt.Printf("i:%d m:%d\n", i, m)
		reverse_num += m * int(math.Pow(2, float64(len(binary)-i-1)))
	}
	return reverse_num
}
func main() {
	num := 43261596
	fmt.Println(Reverse(num))
}
