package main

import "fmt"

func main() {
	fmt.Println(compareSlices([]int{1, 2, 3}, []int{3, 2, 1}))
	fmt.Println(compareSlices([]int{1, 2, 3}, []int{3, 2, 2}))
}
func compareSlices(slice1 []int, slice2 []int) bool {
	if len(slice1) != len(slice2) {
		return false
	}

	elementCount := make(map[int]int)
	for _, num := range slice1 {
		elementCount[num]++
	}

	for _, num := range slice2 {
		elementCount[num]--
		if elementCount[num] < 0 {
			return false
		}
	}

	for _, count := range elementCount {
		if count != 0 {
			return false
		}
	}

	return true
}
 
