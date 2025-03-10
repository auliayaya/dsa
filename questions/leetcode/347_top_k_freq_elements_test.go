package leetcode

import (
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func topKFrequent(nums []int, k int) []int {
	freqMap := make(map[int]int)
	for _, v := range nums {
		freqMap[v]++
	}
	uniqueNums := make([]int, 0, len(freqMap))
	for num := range freqMap {
		uniqueNums = append(uniqueNums, num)
	}
	sort.Slice(uniqueNums, func(i, j int) bool {
		return uniqueNums[i] < uniqueNums[j]
	})
	if k > len(uniqueNums) {
		k = len(uniqueNums)
	}
	return uniqueNums[:k]
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
