package leetcode

import (
	"github.com/stretchr/testify/assert"
	"slices"
	"testing"
)

func sortedSquares(nums []int) []int {
	res := make([]int, len(nums))
	for i, v := range nums {
		res[i] = v * v
	}
	slices.Sort(res)
	return res
}
func TestSortedSquares(t *testing.T) {
	tc := []struct {
		nums   []int
		output []int
	}{
		{
			nums:   []int{-4, -1, 0, 3, 10},
			output: []int{0, 1, 9, 16, 100},
		},
		{
			nums:   []int{-7, -3, 2, 3, 11},
			output: []int{4, 9, 9, 49, 121},
		},
	}
	for _, tt := range tc {
		res := sortedSquares(tt.nums)
		assert.Equal(t, tt.output, res)
	}
}
