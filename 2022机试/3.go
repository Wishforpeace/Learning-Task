package main

import "fmt"

func MaxNum(nums []int) int {
	max := 0
	for _, m := range nums {
		if max < m {
			max = m
		}
	}
	count := make([]int, max+1)
	for _, m := range nums {
		count[m]++
	}
	max_count := 0
	max_num := 0
	for i, m := range count {
		if max_count < m {
			max_count = m
			max_num = i
		}
	}
	return max_num
}
func main() {
	nums := []int{1, 1, 1, 1, 3, 4, 4, 4, 4, 5, 5, 7, 7, 7, 7, 7, 7}
	fmt.Println(MaxNum(nums))
}
