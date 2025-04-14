package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFinNumbers(t *testing.T) {
	tc := []struct {
		nums []int
		out  int
	}{
		{nums: []int{12, 345, 2, 6, 7896}, out: 2},
		{nums: []int{555, 901, 482, 1771}, out: 1},
	}
	for _, tt := range tc {
		res := findNumbers(tt.nums)
		assert.Equal(t, tt.out, res)
	}
}
