package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func trap(height []int) int {
	left, right := 0, len(height)-1
	maxLeft, maxRight := 0, 0
	total := 0
	for left < right {
		maxLeft = max(maxLeft, height[left])
		maxRight = max(maxRight, height[right])
		if height[left] <= height[right] {
			total += min(maxLeft, maxRight) - height[left]
			left++
		} else {
			total += min(maxLeft, maxRight) - height[right]
			right--
		}
	}
	return total
}

func trapStack(height []int) int {
	var stack []int // Stack to store indices
	total := 0      // Total trapped water

	for i := 0; i < len(height); i++ {
		fmt.Println("I ", i)
		// If current height is greater than stack top, process trapped water
		for len(stack) > 0 && height[i] > height[stack[len(stack)-1]] {
			fmt.Println("Stack ", height[i], height[stack[len(stack)-1]])
			bottomIndex := stack[len(stack)-1] // Get the lowest point
			stack = stack[:len(stack)-1]       // Pop it from stack

			if len(stack) == 0 {
				break // No left boundary, so no trapped water
			}

			leftIndex := stack[len(stack)-1] // Left boundary index
			width := i - leftIndex - 1       // Distance between boundaries
			heightDiff := min(height[leftIndex], height[i]) - height[bottomIndex]
			total += width * heightDiff // Water trapped
		}

		stack = append(stack, i) // Push current index onto stack
	}

	return total
}

// de
func TestTrappingRainWater(t *testing.T) {
	tc := []struct {
		height []int
		total  int
	}{
		{
			height: []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1},
			total:  6,
		},
		{
			height: []int{4, 2, 0, 3, 2, 5},
			total:  9,
		},
		{
			height: []int{11, 0, 679},
			total:  11,
		},
		{
			height: []int{11, 679},
			total:  0,
		},
	}
	for _, tt := range tc {
		res := trap(tt.height)
		assert.Equal(t, tt.total, res)
	}
	for _, tt := range tc {
		res := trapStack(tt.height)
		assert.Equal(t, tt.total, res)
	}
}
