package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func removeElement(nums []int, val int) int {
	index := 0
	for _, v := range nums {
		if v != val {
			nums[index] = v
			index += 1
		}
	}
	return index
}

func TestRemoveEleemnt(t *testing.T) {
	tc := []struct {
		input  []int
		k      int
		output int
	}{
		{
			input:  []int{3, 2, 2, 3},
			k:      3,
			output: 2,
		},
		{
			input:  []int{0, 1, 2, 2, 3, 0, 4, 2},
			k:      2,
			output: 5,
		},
	}
	for _, tt := range tc {
		res := removeElement(tt.input, tt.k)
		assert.Equal(t, tt.output, res)
	}
}
