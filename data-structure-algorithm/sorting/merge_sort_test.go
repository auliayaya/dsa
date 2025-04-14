package sorting

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

//Input: nums1 = [1,2,3,0,0,0], m = 3, nums2 = [2,5,6], n = 3
//Output: [1,2,2,3,5,6]
//Explanation: The arrays we are merging are [1,2,3] and [2,5,6].
//The result of the merge is [1,2,2,3,5,6] with the underlined elements coming from nums1.

//Input: nums1 = [1], m = 1, nums2 = [], n = 0
//Output: [1]
//Explanation: The arrays we are merging are [1] and [].
//The result of the merge is [1].

//Input: nums1 = [0], m = 0, nums2 = [1], n = 1
//Output: [1]
//Explanation: The arrays we are merging are [] and [1].
//The result of the merge is [1].
//Note that because m = 0, there are no elements in nums1. The 0 is only there to ensure the merge result can fit in nums1.

func merge(nums1 []int, m int, nums2 []int, n int) {
	copy(nums1[m:], nums2[:n])
	sort.Ints(nums1)
}

func TestMergeSort(t *testing.T) {
	tc := []struct {
		nums1  []int
		nums2  []int
		n      int
		m      int
		output []int
	}{
		{
			nums1:  []int{1, 2, 3, 0, 0, 0},
			nums2:  []int{2, 5, 6},
			m:      3,
			n:      3,
			output: []int{1, 2, 2, 3, 5, 6},
		},
	}
	for _, tt := range tc {
		merge(tt.nums1, tt.m, tt.nums2, tt.n)
		assert.Equal(t, tt.output, tt.nums1)
	}
}
