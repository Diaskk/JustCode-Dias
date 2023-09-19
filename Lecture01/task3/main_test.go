package main

import "testing"

func TestCompareSlices(t *testing.T) {
	testCases := []struct {
		slice1   []int
		slice2   []int
		expected bool
	}{
		{[]int{1, 2, 3}, []int{3, 2, 1}, true},
		{[]int{1, 2, 3}, []int{3, 2, 2}, false},
	}

	for _, testCase := range testCases {
		result := compareSlices(testCase.slice1, testCase.slice2)
		if result != testCase.expected {
			t.Errorf("compareSlices(%v, %v) = %v; want %v", testCase.slice1, testCase.slice2, result, testCase.expected)
		}
	}
}
