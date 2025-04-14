package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func isAnagram(s string, t string) bool {
	str := strings.ToLower(s)
	str2 := strings.ToLower(t)
	charCount := make(map[rune]int)
	if len(str2) != len(str) {
		return false
	}
	for _, v := range str {
		charCount[v]++
	}
	for _, v := range str2 {
		charCount[v]--
		if charCount[v] < 0 {
			return false
		}
	}
	return true
}
func Test_IsAnagram(t *testing.T) {
	v := []struct {
		data   []string
		result bool
	}{
		{
			data:   []string{"test", "tset"},
			result: true,
		},
		{
			data:   []string{"anagram", "nagaram"},
			result: true,
		},
		{
			data:   []string{"rat", "car"},
			result: false,
		},
		{
			data:   []string{"Cheater", "teacher"},
			result: true,
		},
		{
			data:   []string{"God", "dog"},
			result: true,
		},
		{
			data:   []string{"Planter", "replant"},
			result: true,
		},
		{
			data:   []string{"ayas", "udin"},
			result: false,
		},
	}
	for _, test := range v {
		rt := isAnagram(test.data[0], test.data[1])
		assert.Equal(t, rt, test.result)
	}
}
