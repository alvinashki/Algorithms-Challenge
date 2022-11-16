package main

import "fmt"

func breakingRecords(scores []int32) []int32 {
	var min, max, tempmin, tempmax int32
	result := []int32{}
	min, max = scores[0], scores[0]
	for i := 1; i < len(scores); i++ {
		if scores[i] < min {
			min = scores[i]
			tempmin++
		}
		if scores[i] > max {
			max = scores[i]
			tempmax++
		}
	}
	result = append(result, tempmax, tempmin)
	return result
}

func main() {
	fmt.Println(breakingRecords([]int32{10, 5, 20, 20, 4, 5, 2, 25, 1}))
	fmt.Println(breakingRecords([]int32{3, 4, 21, 36, 10, 28, 35, 5, 24, 42}))
	fmt.Println(breakingRecords([]int32{10, 12, 4, 5, 9}))
}
