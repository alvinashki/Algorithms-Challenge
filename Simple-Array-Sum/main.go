package main

import "fmt"

func simpleArraySum(ar []int32) int32 {
	var result int32
	for _, v := range ar {
		result = result + v
	}
	return result
}

func main() {
	fmt.Println(simpleArraySum([]int32{1, 2, 3, 4, 10, 11}))
	fmt.Println(simpleArraySum([]int32{5, 2, 6, 1, 2}))
}
