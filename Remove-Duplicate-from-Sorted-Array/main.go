package main

import "fmt"

func removeDuplicates(nums []int) int {
	slice := []int{}
	for i := 0; i < len(nums)-1; i++ {
		j := i + 1
		k := i - 1

		if i == 0 {
			if nums[i] == nums[j] {
				slice = append(slice, nums[j])
			} else {
				slice = append(slice, nums[i])
			}
		}
		if i > 0 && i <= len(nums)-3 {
			if nums[i] != nums[j] && nums[i] != nums[k] {
				slice = append(slice, nums[i])
			}
			if nums[i] == nums[j] && nums[i] != nums[k] {
				slice = append(slice, nums[i])
			}
		}
		if i == len(nums)-2 {
			if nums[i] != nums[j] {
				slice = append(slice, nums[j])
			} else {
				slice = append(slice, nums[j])
			}
		}
	}
	result := len(slice)
	fmt.Println("result =", slice)

	return result
}

func main() {
	fmt.Println(removeDuplicates([]int{1, 1, 2, 3, 3, 3, 9}))
	fmt.Println(removeDuplicates([]int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}))
	fmt.Println(removeDuplicates([]int{1, 1, 2}))
}
