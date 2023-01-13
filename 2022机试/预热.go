package main

import "fmt"

func EligibleNum(number int, nums []int) int {
	var count = 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == number {
				count++
			}｜
		}
	}
	return count
}
func main() {
	var number = 11
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(EligibleNum(number, nums))
}
