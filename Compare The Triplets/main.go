package main

import "fmt"

func compareTriplets(a []int32, b []int32) []int32 {
	slice := []int32{}
	var tempa, tempb int
	for i := 0; i < len(a); i++ {
		if a[i] > b[i] {
			tempa += 1
		} else if a[i] < b[i] {
			tempb += 1
		}
	}
	slice = append(slice, int32(tempa), int32(tempb))
	return slice
}

func main() {
	fmt.Println(compareTriplets([]int32{1, 2, 3}, []int32{3, 2, 1}))
	fmt.Println(compareTriplets([]int32{5, 6, 7}, []int32{3, 6, 10}))
	fmt.Println(compareTriplets([]int32{17, 28, 30}, []int32{99, 16, 8}))
}
