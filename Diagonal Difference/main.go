package main

import "fmt"

func diagonalDifference(arr [][]int32) int32 {
	var a, b, result int32

	k := len(arr) - 1
	l := k - 1
	for i := 0; i < len(arr)-1; i++ {
		j := i + 1

		if i == 0 {
			a = a + arr[i][i] + arr[j][j]
			b = b + arr[i][k] + arr[j][l]
		} else if i > 0 {
			a = a + arr[j][j]
			b = b + arr[j][l]
		}
		j++
		l--
	}

	if a > b {
		result = a - b
	} else {
		result = b - a
	}
	return result
}

func main() {
	fmt.Println(diagonalDifference([][]int32{{1, 2, 3}, {3, 2, 1}, {6, 5, 2}}))
	fmt.Println(diagonalDifference([][]int32{{1, 2, 3}, {4, 5, 6}, {9, 8, 7}}))
}
