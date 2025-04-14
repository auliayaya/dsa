package leetcode

import "testing"

func rotate(nums []int, k int) {
	n := len(nums)
	k %= n
	reverse(nums, 0, n-1)
	reverse(nums, 0, k-1)
	reverse(nums, k, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		temp := nums[start]
		nums[start] = nums[end]
		nums[end] = temp
		start++
		end--
	}
}

// Steps Explanation:
// Reverse the entire array:
// Example: ( [1,2,3,4,5,6,7] ) → ( [7,6,5,4,3,2,1] )
// Reverse the first ( k ) elements:
// Example: ( [7,6,5,4,3,2,1] ) → ( [5,6,7,4,3,2,1] )
// Reverse the last ( n-k ) elements:
// Example: ( [5,6,7,4,3,2,1] ) → ( [5,6,7,1,2,3,4] )
func TestRotateArray(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
		k      int
	}{
		{
			input:  []int{1, 2, 3, 4, 5, 6, 7},
			output: []int{5, 6, 7, 1, 2, 3, 4},
			k:      3,
		},
		{
			input:  []int{-1, -100, 3, 99},
			output: []int{3, 99, -1, -100},
			k:      2,
		},
	}
	for _, tt := range tc {
		rotate(tt.input, tt.k)

	}
}
