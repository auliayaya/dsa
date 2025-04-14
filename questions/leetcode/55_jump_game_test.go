package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func canJump(nums []int) bool {
	maxReach := 0
	for i, v := range nums {
		if i > maxReach {
			return false
		}
		maxReach = max(maxReach, v+i)
		if maxReach >= len(nums)-1 {
			return true
		}
	}
	return false
}

func TestJumpGame(t *testing.T) {
	tc := []struct {
		input []int
		out   bool
	}{
		{
			input: []int{2, 3, 1, 1, 4},
			out:   true,
		},
		{
			input: []int{3, 2, 1, 0, 4},
			out:   false,
		},
	}
	for _, tt := range tc {
		res := canJump(tt.input)
		assert.Equal(t, tt.out, res)
	}
}
