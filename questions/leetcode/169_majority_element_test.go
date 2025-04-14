package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func majorityElement(nums []int) int {
	mp := map[int]int{}
	for _, v := range nums {
		mp[v]++
		if mp[v] > len(nums)/2 {
			return v
		}
	}
	return -1
}

// Optimized Approach: Boyer-Moore Voting Algorithm
// Since the problem guarantees that a majority element always exists, we can find it in O(n) time and O(1) space.

// Time & Space Complexity
// Approach	Time Complexity	Space Complexity
// Hash Map	O(n)	O(n)
// Boyer-Moore	O(n)	O(1)
// 🔹 Use Boyer-Moore when the majority element is guaranteed to exist.
// 🔹 Use Hash Map if you need to count all elements (for variation problems).
func majorityElementOptimized(nums []int) int {
	count, candidate := 0, 0

	for _, v := range nums {
		if count == 0 {
			candidate = v
		}
		if v == candidate {
			count++
		} else {
			count--
		}

	}
	return candidate
}

func TestMajorityElement(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 3},
			output: 3,
		},
		{
			input:  []int{2, 2, 1, 1, 1, 2, 2},
			output: 2,
		},
	}
	for _, tt := range tc {
		actual := majorityElement(tt.input)
		assert.Equal(t, tt.output, actual)
	}
}

func TestMajorityElementOptimized(t *testing.T) {
	tc := []struct {
		input  []int
		output int
	}{
		{
			input:  []int{3, 2, 3},
			output: 3,
		},
		{
			input:  []int{2, 2, 1, 1, 1, 2, 2},
			output: 2,
		},
	}
	for _, tt := range tc {
		actual := majorityElementOptimized(tt.input)
		assert.Equal(t, tt.output, actual)
	}
}
