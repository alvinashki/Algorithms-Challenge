package main

import (
	"fmt"
	"strconv"
)

func main() {
	slice := []string{}
	a, b := 3, 5
	c := a * b
	for i := 1; i <= 100; i++ {
		if i == c {
			slice = append(slice, "snip-snap")
			c = c + 15
			a = a + 3
			b = b + 5
		} else if i == a {
			slice = append(slice, "snip")
			a = a + 3
		} else if i == b {
			slice = append(slice, "snap")
			b = b + 5
		} else {
			s := strconv.Itoa(i)
			slice = append(slice, string(s))
		}
	}
	fmt.Println(slice)
}
