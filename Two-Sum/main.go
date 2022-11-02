package main

import (
	"fmt"
)

func twoSum(nums []int, target int) []int {
	var slice = []int{}

	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			k := j
			temp := nums[i] + nums[k]
			if temp == target {
				slice = append(slice, i, k)
			}
			k++
		}
	}
	return slice
}

func main() {
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 9))
	fmt.Println(twoSum([]int{3, 2, 4}, 6))
	fmt.Println(twoSum([]int{2, 7, 11, 15}, 26))
}
