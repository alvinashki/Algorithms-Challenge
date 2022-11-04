package main

import "fmt"

func birthday(s []int32, d, m int32) int32 {
	var result, sum int32

	for i := 0; i < len(s); i++ {
		var temp int32
		sum = s[i]
		if m == 1 {
			if s[i] == d {
				result++
				break
			}
		}
		for j := i + 1; j < len(s); j++ {
			temp++
			sum = sum + s[j]
			if temp == m-1 {
				if sum == d {
					result++
				}
				j = len(s)
			}
		}
	}
	return result
}

func main() {
	fmt.Print(birthday([]int32{1, 4}, 4, 1))
	fmt.Print(" ", birthday([]int32{1, 2, 1, 3, 2}, 3, 2))
	fmt.Print(" ", birthday([]int32{2, 5, 1, 3, 4, 4, 3, 5, 1, 1, 2, 1, 4, 1, 3, 3, 4, 2, 1}, 18, 7))
	fmt.Print(" ", birthday([]int32{2, 2, 2, 1, 3, 2, 2, 3, 3, 1, 4, 1, 3, 2, 2, 1, 2, 2, 4, 2, 2, 3, 5, 3, 4, 3, 2, 1, 4, 5, 4}, 10, 4))
	fmt.Println()
}
