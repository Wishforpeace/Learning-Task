package main

import (
	"fmt"
	"math"
)

func main() {
	var array []int
	var MaxNum = 0.0
	var max int
	var num int
	for {
		fmt.Scanf("%d", &num)
		if num == 0 {
			break
		}
		array = append(array, num)
	}
	for i, m := range array {
		abs := math.Abs(float64(m))
		if abs > MaxNum {
			MaxNum = abs
			max = i
		}
		//println(m)
	}
	fmt.Println("min:", array[max])
}
