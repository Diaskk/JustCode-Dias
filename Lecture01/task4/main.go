package main

import (
	"fmt"
	"sort"
)

func main() {
	s := intSlice{slice: []int{4, 3, 2, 5, 1}}
	s.sort()
	fmt.Println(s.slice)
}

type intSlice struct {
	slice []int
}

func (is *intSlice) sort() {
	sort.Ints(is.slice)
}
