package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {

	var result int64
	for i := 0; i < len(ar); i++ {
		result += ar[i]
	}
	return result

}

func main() {
	fmt.Println(aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005}))
}
