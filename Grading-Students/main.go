package main

import "fmt"

func main() {
	grades := []int32{73, 67, 38, 33}
	compare := []int32{35, 40, 45, 50, 55, 60, 65, 70, 75, 80, 85, 90, 95, 100}
	result := []int32{}
	var temp, diff int32

	for _, vG := range grades {
		if vG > 35 {
			for i := 0; compare[i] < vG; i++ {
				temp = compare[i] + 5
			}
			diff = temp - vG
			if diff < 3 {
				result = append(result, temp)
			} else {
				result = append(result, vG)
			}
		} else {
			result = append(result, vG)
		}
	}
	fmt.Println(result)
}
