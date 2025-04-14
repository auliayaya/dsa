package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxProfit2(prices []int) int {
	maxProf := 0
	j := 1
	for i := 0; i < len(prices)-1; i++ {
		result := prices[j] - prices[i]
		if result > 0 {
			maxProf += result
		}
		j++
	}
	return maxProf
}

func TestMaxProfit2(t *testing.T) {
	tc := []struct {
		input  []int
		expect int
	}{
		{
			input:  []int{7, 1, 5, 3, 6, 4},
			expect: 7,
		},
		{
			input:  []int{1, 2, 3, 4, 5},
			expect: 4,
		},
		{
			input:  []int{7, 6, 4, 3, 1},
			expect: 0,
		},
	}
	for _, tt := range tc {
		actual := maxProfit2(tt.input)
		assert.Equal(t, tt.expect, actual)
	}
}
