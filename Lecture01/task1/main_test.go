package main

import "testing"

func TestTwoSum(t *testing.T) {
	nums := []int{2, 7, 11, 15}
	target := 9
	result := twoSum(nums, target)
	expected := []int{0, 1}

	if !equalSlices(result, expected) {
		t.Errorf("Expected %v, but got %v", expected, result)
	}
}
func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}
