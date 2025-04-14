package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeDuplicates(nums []int) int {
	if len(nums) <= 2 {
		return len(nums)
	}
	j := 2
	for i := 2; i < len(nums); i++ {
		if nums[i] != nums[j-2] {
			nums[j] = nums[i]
			j++
		}
	}
	return j
}

func TestRemoveDuplicates(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{1, 1, 1, 2, 2, 3},
			output: 5,
		},
		{
			input:  []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			output: 7,
		},
	}
	for _, tt := range tc {
		actual := removeDuplicates(tt.input)
		assert.Equal(t, tt.output, actual)
	}
}
