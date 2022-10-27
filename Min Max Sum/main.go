package main

import "fmt"

func miniMaxSum(arr []int32) {
	slice := []int{}

	for i := 0; i < len(arr); i++ {
		temp := 0
		for j := 0; j < len(arr); j++ {
			if j != i {
				temp += int(arr[j])
			}

		}
		slice = append(slice, temp)
	}

	minval, maxval := slice[0], slice[0]
	for _, val := range slice {
		if val > maxval {
			maxval = val
		}
		if val < minval {
			minval = val
		}
	}

	fmt.Println(minval, maxval)

}

func main() {
	miniMaxSum([]int32{1, 3, 5, 7, 9})
	miniMaxSum([]int32{1, 2, 3, 4, 5})
}
