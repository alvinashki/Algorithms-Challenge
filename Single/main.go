package main

import "fmt"

func single(array []int) {
	for i, v := range array {
		var temp int
		for _, x := range array {
			if v == x {
				temp++
			}
		}
		if temp < 2 {
			fmt.Println(v, i+1)
		}
	}
}

func main() {
	single([]int{102, 32, 99, 32, 45, 102, 45, 67, 67, 100, 100})
	single([]int{101, 32, 93, 90, 32, 65, 101, 90, 93, 1})
}
