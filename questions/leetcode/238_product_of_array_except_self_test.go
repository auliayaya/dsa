package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func productExceptSelf(nums []int) []int {
	n := len(nums)
	answer := make([]int, n)
	prefix := 1
	suffix := 1
	for i := 0; i < n; i++ {
		answer[i] = prefix
		prefix *= nums[i]
	}
	for i := n - 1; i >= 0; i-- {
		answer[i] *= suffix
		suffix *= nums[i]
	}
	return answer
}

func TestProductExceptSelf(t *testing.T) {
	tc := []struct {
		nums   []int
		expect []int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			expect: []int{24, 12, 8, 6},
		},
	}
	for _, tt := range tc {
		res := productExceptSelf(tt.nums)
		assert.Equal(t, tt.expect, res)
	}
}
