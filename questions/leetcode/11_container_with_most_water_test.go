package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxArea(height []int) int {
	left, right, maxArea := 0, len(height)-1, 0

	for left < right {
		currentArea := min(height[left], height[right]) * (right - left)
		maxArea = max(maxArea, currentArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}
func TestMaxArea(t *testing.T) {
	tc := []struct {
		height []int
		ans    int
	}{
		{
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			ans:    49,
		}, {
			height: []int{1, 1},
			ans:    1,
		},
	}
	for _, tt := range tc {
		maxArea := maxArea(tt.height)
		assert.Equal(t, tt.ans, maxArea)
	}
}
