package leetcode

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func isAnagram(s string, t string) bool {
	if s == "" || t == "" {
		return false
	}
	if s == t {
		return true
	}
	var str, str2 []string
	for _, v := range s {
		str = append(str, string(v))
	}
	for _, v := range t {
		str = append(str2, string(v))
	}

	if strings.Join(str, "") == strings.Join(str2, "") {
		return true
	}
	return false
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
			result: true,
		},
		{
			data:   []string{"Cheater", "teacher"},
			result: true,
		},
		{
			data: []string{"God","dog"},
			result: true,
		},
		{
			data : []string{"Planter", "replant"},
		},
		Cheater = teacher.
		God = dog.
		Planter = replant.
		Lampshade = headlamps.
		Bust = stub.
		Roots = torso.
		Rail = lair.
		Donate = atoned.
	}
	for _, test := range v {
		rt := isAnagram(test.data[0], test.data[1])
		assert.Equal(t, rt, test.result)

	}
}
