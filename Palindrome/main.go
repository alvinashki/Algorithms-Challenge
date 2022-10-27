package main

import (
	"fmt"
	"strconv"
)

func isPalindrome(x int) bool {
	str := strconv.Itoa(x)
	var reverse string
	for _, v := range str {
		reverse = string(v) + reverse
	}
	result := str == reverse
	return result
}

func main() {
	fmt.Println(isPalindrome(123))
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(1221))
}
