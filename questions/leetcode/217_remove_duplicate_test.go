package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func containsDuplicate(nums []int) bool {
	abc := map[int]bool{}
	for _, v := range nums {
		if _, ok := abc[v]; ok {
			return true
		}
		abc[v] = true
	}
	return false
}

func TestContainsDuplicate(t *testing.T) {
	vals := [][]int{
		{
			1, 3, 4, 5, 5,
		}, {
			1, 2, 3, 1,
		},
		{
			1, 1, 1, 3, 3, 4, 3, 2, 4, 2,
		},
	}
	for _, v := range vals {
		result := containsDuplicate(v)
		assert.Equal(t, true, result)
	}
}
