package main

import "fmt"

func birthdayCakeCandles(candles []int32) int32 {
	var result, value int32

	for _, value1 := range candles {
		if value1 > value {
			value = value1
		}
	}

	for _, val2 := range candles {
		if val2 == value {
			result += 1
		}
	}

	return result
}

func main() {
	fmt.Println(birthdayCakeCandles([]int32{3, 2, 1, 3}))
	fmt.Println(birthdayCakeCandles([]int32{2, 6, 3, 4, 3}))
}
