package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

func sorting(s string) string {
	rn := []rune(s)
	sort.Slice(rn, func(i, j int) bool {
		return rn[i] < rn[j]
	})
	return string(rn)
}

func groupAnagrams(strs []string) [][]string {
	anagrams := make(map[string][]string, 0)
	for i, _ := range strs {
		fmt.Println(anagrams[sorting(strs[i])])
		sortedWord := sorting(strs[i])
		anagrams[sortedWord] = append(anagrams[sortedWord], strs[i])
	}
	res := make([][]string, 0, len(anagrams))
	for i, _ := range anagrams {
		res = append(res, anagrams[i])
	}
	return res
}

func TestGroupOfAnagrams(t *testing.T) {
	tc := []struct {
		strs   []string
		result [][]string
	}{
		{
			strs:   []string{"eat", "tea", "tan", "ate", "nat", "bat"},
			result: [][]string{{"bat"}, {"nat", "tan"}, {"ate", "eat", "tea"}},
		},
		{
			strs:   []string{""},
			result: [][]string{{""}},
		},
		{
			strs:   []string{"a"},
			result: [][]string{{"a"}},
		},
	}
	for _, tt := range tc {
		actual := groupAnagrams(tt.strs)
		assert.Equal(t, tt.result, actual)
	}
}
