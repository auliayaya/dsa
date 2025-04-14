package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}
	uniqueNums := make([][]int, len(nums)+1)
	for num, freq := range freqMap {
		uniqueNums[freq] = append(uniqueNums[freq], num)
	}
	var res []int
	for i := len(nums); i >= 0 && len(res) < k; i-- {
		if len(uniqueNums[i]) > 0 {
			res = append(res, uniqueNums[i]...)
		}
	}
	return res[:k]
}

func TestTopKFrequent(t *testing.T) {
	tc := []struct {
		input  []int
		result []int
		k      int
	}{
		{
			input:  []int{1, 1, 1, 2, 2, 3},
			k:      2,
			result: []int{1, 2},
		},
		{
			input:  []int{1, 1, 1, 2, 2, 3, 4, 4, 4},
			k:      3,
			result: []int{1, 4, 2},
		},
		{
			input:  []int{1},
			k:      1,
			result: []int{1},
		},
	}
	for _, tt := range tc {
		result := topKFrequent(tt.input, tt.k)
		assert.Equal(t, tt.result, result)
	}
}
