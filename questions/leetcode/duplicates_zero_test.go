package leetcode

import "testing"

func duplicateZeros(arr []int) {

}

func TestDuplicateZeros(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 0, 2, 3, 0, 4, 5, 0}, output: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{input: []int{}, output: []int{}},
	}
}
