package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRemoveDuplicatesWithStack(t *testing.T) {
	want := "ca"
	got := removeDuplicateWithStack("abbaca")
	assert.Equal(t, want, got)
}

func TestRemoveDuplicatesWithPointer(t *testing.T) {
	want := "ca"
	got := removeDuplicateTwoPointer("abbaca")
	assert.Equal(t, want, got)
}
func BenchmarkRemoveDuplicateWithPointer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		removeDuplicateTwoPointer("abbaca")
	}
}
