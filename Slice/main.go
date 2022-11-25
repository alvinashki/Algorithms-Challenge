package main

import "fmt"

func function(x, y []int) []int {
	slice := []int{}

	for i := 0; i < len(x); i++ {
		if i == 0 {
			for j := 0; j < len(y); j++ {
				if x[i] == y[j] {
					slice = append(slice, x[i])
					j = len(y)
				}
			}
		}
		if i > 0 {
			if x[i] != x[i-1] {
				for j := 0; j < len(y); j++ {
					if x[i] == y[j] {
						slice = append(slice, x[i])
						j = len(y)
					}
				}
			}
		}
	}

	return slice
}

func main() {
	fmt.Println(function([]int{1, 2, 3, 3, 3, 4, 4, 5}, []int{2, 3, 4, 5}))
	fmt.Println(function([]int{2, 3, 4, 5}, []int{1, 2, 3, 3, 3, 4, 4, 5}))
	fmt.Println(function([]int{1,4,5,5,7,9,9}, []int{2,2,3,4,4,5,9}))
}
