package main

import "fmt"

func staircase(n int32) {
	for i := 0; i < int(n); i++ {
		for j := int(n) - i - 1; j > 0; j-- {
			fmt.Print(" ")
		}

		for k := 0; k <= i; k++ {
			fmt.Print("#")
		}
		fmt.Println()
	}
}

func main() {
	staircase(3)
	staircase(5)
}
