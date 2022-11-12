package main

import (
	"fmt"
	"sort"
)

func arrayChallenge(arr []int) bool {
	sort.Ints(arr)
	var result bool

	for i := 0; i < len(arr)-2; i++ {
		var temp int
		k := i + 1
		for j := 0; j < len(arr)-2; j++ {
			if arr[i] != arr[j] {
				temp = temp + arr[i] + arr[j]
				fmt.Println("temp 1 = ", temp)
				if temp == arr[len(arr)-1] {
					result = true
					return result
				}

			}
		}
		for j := 0; j < len(arr)-2; j++ {
			if arr[i] != arr[j] && j != k {
				temp = temp + arr[i] + arr[j]
				fmt.Println("temp 2 = ", temp)
				if temp == arr[len(arr)-1] {
					result = true
					return result
				}
			}
		}
		k++
	}
	return result
}

func main() {
	fmt.Println(arrayChallenge([]int{3, 5, -1, 8, 12}))
	fmt.Println(arrayChallenge([]int{7, 3, 9, 6, 5, 19}))
}
