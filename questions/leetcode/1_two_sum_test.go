package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func twoSum(nums []int, target int) []int {
	var idx []int
	temp := map[int]int{}
	for i, v := range nums {
		if k, ok := temp[target-v]; ok {
			return []int{k, i}
		}
		temp[v] = i
	}
	// target 3
	// 3-1 => return
	// 1=0
	// 3-2 => return 0,1
	// target 7
	// 7-1 return
	// 1=0
	// 7-1 => return
	// 1=0
	// 7-2 => return
	// 2=1
	// 7-3 => return
	// 3=2
	// 7-4 => return k, i
	return idx
}

func TestTwoSum(t *testing.T) {
	tc := []struct {
		nums   []int
		result []int
		target int
	}{
		{
			nums:   []int{1, 2, 3, 4},
			result: []int{0, 1},
			target: 3,
		},
		{
			nums:   []int{1, 2, 3, 4},
			result: []int{2, 3},
			target: 7,
		},
		{
			nums:   []int{1, 2, 3, 4},
			result: []int{1, 3},
			target: 6,
		},
	}
	for _, tt := range tc {
		got := twoSum(tt.nums, tt.target)
		fmt.Println("Idx", got)
		assert.Equal(t, tt.result, got)
	}
}
