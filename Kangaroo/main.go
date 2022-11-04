package main

import (
	"fmt"
)

func kangaroo(x1 int32, v1 int32, x2 int32, v2 int32) string {
	temp1, temp2 := x1+v1, x2+v2
	temp := temp1 * temp2
	var i int32
	s := "NO"

	for i = 0; i <= temp; i++ {
		if temp1 == temp2 {
			s = "YES"
			break
		}
		temp1 = temp1 + v1
		temp2 = temp2 + v2
	}
	return s

}

func main() {
	fmt.Println(kangaroo(0, 3, 4, 2))
	fmt.Println(kangaroo(2, 1, 1, 2))
}
