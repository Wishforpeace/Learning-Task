package main

import (
	"fmt"
	"math"
	"reflect"
	"strconv"
)

func Solve(nums []int, sum int) (int, int) {
	choice := int(math.Pow(2, 10))
	cout := 0
	max := 1
	var final [][]int
	var final_tmp [][]int
	for i := 0; choice < int(math.Pow(2, 11))-1; i++ {
		choice = choice + 1
		max_tmp := 1
		binary_choice := Binary(choice)
		//fmt.Println(binary_choice)
		vol := 0
		var result []int
		for j := 1; j < len(binary_choice); j++ {
			option, _ := strconv.Atoi(string(binary_choice[j]))
			//fmt.Println(option)
			vol += option * nums[j-1]
			if option == 1 {
				max_tmp = max_tmp * nums[j-1]
				result = append(result, nums[j-1])
			}
			//fmt.Println("max", max_tmp)
		}
		if vol == sum {
			//fmt.Println(result)
			final_tmp = append(final_tmp, result)
			if max_tmp > max {
				max = max_tmp
			}
		}
	}
	final = DeleteRepeated(final_tmp)
	//fmt.Println(final_tmp)
	fmt.Println(final)
	cout = len(final)
	return cout, max
}
func Binary(number int) string {
	var binary string
	for number != 0 {
		//fmt.Println(number % 2)
		binary = strconv.Itoa(number%2) + binary
		number = number / 2
	}
	return binary
}
func DeleteRepeated(nums [][]int) [][]int {
	newRes := make([][]int, 0)
	for i := 0; i < len(nums); i++ {
		flag := false
		for j := i + 1; j < len(nums); j++ {
			if reflect.DeepEqual(nums[i], nums[j]) {
				flag = true
				break
			}
		}
		if !flag {
			newRes = append(newRes, nums[i])
		}
	}
	return newRes

}
func main() {
	//n := 10
	N := 6
	nums := []int{1, 2, 4, 1, 3, 3, 5, 2, 6, 2}
	count, product := Solve(nums, N)
	fmt.Println(count, product)
}
