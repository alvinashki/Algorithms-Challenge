package main

import "fmt"

func mainn(s, t, a, b int32, apples, oranges []int32) []int32 {
	tempA, tempO := []int32{}, []int32{}
	var app, orr int32

	for _, v := range apples {
		tempA = append(tempA, v+a)
	}
	for _, v := range oranges {
		tempO = append(tempO, v+b)
	}
	for _, v := range tempA {
		if v >= s {
			if v <= t {
				app = app + 1
			}
		}
	}
	for _, v := range tempO {
		if v >= b-(b-s) {
			if v <= b-(b-t) {
				orr = orr + 1
			}
		}
	}
	slice := []int32{}
	slice = append(slice, app, orr)
	return slice
}

func main() {
	fmt.Println(mainn(7, 10, 4, 12, []int32{2, 3, -4}, []int32{3, -2, -4}))
	fmt.Println(mainn(7, 11, 5, 15, []int32{-2, 2, 1}, []int32{5, -6}))
}
