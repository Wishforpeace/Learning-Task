package main

import "fmt"

func PowNum(number int, nums []int) []int {
	for i := 0; i < number; i++ {
		nums[i] = nums[i] * nums[i]
	}
	for i := 0; i < number-1; i++ {
		for j := 0; j < number-i-1; j++ {
			if nums[j] > nums[j+1] {
				temp := nums[i]
				nums[i] = nums[j]
				nums[j] = temp
			}
		}
	}
	return nums
}
func main() {
	var nums = []int{-6, 4, 5, 7, 9}
	pow_num := PowNum(5, nums)
	fmt.Println("平方排序后", pow_num)
}
