package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func maxProfit(prices []int) int {
	maxProf := 0
	minPrice := prices[0]
	n := len(prices)
	for i := 1; i < n; i++ {
		if minPrice > prices[i] {
			minPrice = prices[i]
		} else {
			maxProf = max(maxProf, prices[i]-minPrice)
		}
	}
	return maxProf
}

//Time complexity O(n)
// Space complexity O()

func TestMaxProfit(t *testing.T) {
	tc := []struct {
		prices []int
		profit int
	}{
		{
			prices: []int{7, 1, 5, 3, 6, 4},
			profit: 5,
		},
		{
			prices: []int{7, 6, 4, 3, 1},
			profit: 0,
		},
		{
			prices: []int{1, 2},
			profit: 1,
		},
	}
	for _, tt := range tc {
		ast := maxProfit(tt.prices)
		assert.Equal(t, tt.profit, ast)
	}
}
