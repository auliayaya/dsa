package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func findMaxConsecutiveOnes(nums []int) int {
	count := 0
	maxCount := 0
	for _, n := range nums {
		if n == 1 {
			count++
			maxCount = max(maxCount, count)
		} else {
			count = 0
		}
	}
	return maxCount
}

func TestFindMaxConsecutiveOnes(t *testing.T) {
	tc := []struct {
		nums []int
		out  int
	}{
		{
			nums: []int{1, 1, 0, 1, 1, 1},
			out:  3,
		},
		{
			nums: []int{1, 0, 1, 1, 0, 1},
			out:  2,
		},
	}
	for _, tt := range tc {
		res := findMaxConsecutiveOnes(tt.nums)
		assert.Equal(t, tt.out, res)
	}
}
