package main

import "fmt"

func plusMinus(arr []int32) {
	var ratio [3]float64

	for i := 0; i < len(arr); i++ {
		if arr[i] > 0 {
			ratio[0] += 1
		} else if arr[i] < 0 {
			ratio[1] += 1
		} else {
			ratio[2] += 1
		}
	}

	for _, val_r := range ratio {
		fmt.Println(val_r / float64(len(arr)))
	}
}

func main() {
	plusMinus([]int32{-4, 3, -9, 0, 4, 1})
	plusMinus([]int32{1, 2, 3, -1, -2, -3, 0, 0})
}
