package main

import "fmt"

func challenge(arr []int) int {
	var a, b, temp, result int
	for i := 0; i < len(arr)-1; i++ {
		if i == 0 {
			if arr[i] < arr[i+1] {
				a = arr[i]
			} else {
				a = arr[i+1]
			}
		} else if i < len(arr) {
			if a > arr[i] {
				a = arr[i]
				temp = i
			}
		}
	}

	for j := temp; j < len(arr); j++ {
		if j == temp {
			if arr[j] > arr[j+1] {
				b = arr[j]
			} else {
				b = arr[j+1]
			}
		} else {
			if b < arr[j] {
				b = arr[j]
			}
		}

	}

	fmt.Println("A =", a, " B =", b)
	if b <= a {
		return -1
	}
	result = b - a
	return result
}

func main() {
	fmt.Println(challenge([]int{44, 30, 24, 32, 35, 30, 40, 38, 15}))
	fmt.Println(challenge([]int{10, 9, 8, 2}))
	fmt.Println(challenge([]int{10, 12, 4, 5, 9}))
}
