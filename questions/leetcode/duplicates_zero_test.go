package leetcode

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// Brute Force	O(N²)	O(1)	Too slow for large N
// Optimized Two-Pointer (Best Solution)	O(N)	O(1)	Efficient in-place method
// Using Extra Slice	O(N)	O(N)	Simple but uses extra memory
func duplicateZerosOn2(arr []int) {
	n := len(arr)
	for i := 0; i < n; i++ {
		if arr[i] == 0 {
			// Shift all elements to the right
			for j := n - 1; j > i; j-- {
				arr[j] = arr[j-1]
			}
			i++ // Skip the next zero to avoid re-duplication
		}
	}
}
func duplicateZerosExtraSpace(arr []int) {
	n := len(arr)
	newArr := make([]int, n)
	j := 0

	for i := 0; i < n && j < n; i++ {
		newArr[j] = arr[i]
		j++
		if arr[i] == 0 && j < n {
			newArr[j] = 0
			j++
		}
	}

	// Copy back to the original array
	copy(arr, newArr)
}
func duplicateZeros(arr []int) {
	n := len(arr)
	count := 0

	// Count zeros to determine shifting
	for _, num := range arr {
		if num == 0 {
			count++
		}
	}

	// Traverse from the end to shift elements
	i := n - 1
	j := n + count - 1
	//fmt.Println("Count ", count)

	for i >= 0 {
		if j < n {
			arr[j] = arr[i]
			//fmt.Println("Arr ", arr, arr[j], arr[i])
		}
		j--

		// Duplicate zero if found
		if arr[i] == 0 {
			if j < n {
				arr[j] = 0
				//fmt.Println("Arry ", arr, j, arr[j])
			}
			j--
		}
		i--
	}
}

func TestDuplicateZeros(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 0, 2, 3, 0, 4, 5, 0}, output: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{input: []int{1, 2, 3}, output: []int{1, 2, 3}},
	}
	for _, tt := range tc {
		duplicateZeros(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}
func TestDuplicateZerosOn2(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 0, 2, 3, 0, 4, 5, 0}, output: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{input: []int{1, 2, 3}, output: []int{1, 2, 3}},
	}
	for _, tt := range tc {
		duplicateZerosOn2(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}
func TestDuplicateZerosExtraSpace(t *testing.T) {
	tc := []struct {
		input  []int
		output []int
	}{
		{input: []int{1, 0, 2, 3, 0, 4, 5, 0}, output: []int{1, 0, 0, 2, 3, 0, 0, 4}},
		{input: []int{1, 2, 3}, output: []int{1, 2, 3}},
	}
	for _, tt := range tc {
		duplicateZerosExtraSpace(tt.input)
		assert.Equal(t, tt.output, tt.input)
	}
}

var testArr = []int{1, 0, 2, 3, 0, 4, 5, 0} // Sample input

// Benchmark for O(N²) Brute Force
//func BenchmarkDuplicateZerosOn2(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		arr := append([]int{}, testArr...) // Copy to avoid mutation
//		duplicateZerosOn2(arr)
//	}
//}
//
//// Benchmark for O(N) Optimized Two-Pointer Solution
//func BenchmarkDuplicateZeros(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		arr := append([]int{}, testArr...)
//		duplicateZeros(arr)
//	}
//}
//
//// Benchmark for O(N) Extra Space Solution
//func BenchmarkDuplicateZerosExtraSpace(b *testing.B) {
//	for i := 0; i < b.N; i++ {
//		arr := append([]int{}, testArr...)
//		duplicateZerosExtraSpace(arr)
//	}
//}

func BenchmarkDuplicateZerosOn2(b *testing.B) {
	arr := make([]int, 10000) // Large input size (10,000 elements)
	for i := 0; i < len(arr); i += 2 {
		arr[i] = 0 // Insert zeros at even indices
	}

	b.ResetTimer() // Start timing after initialization
	for i := 0; i < b.N; i++ {
		duplicateZerosOn2(arr)
	}
}

func BenchmarkDuplicateZeros(b *testing.B) {
	arr := make([]int, 10000)
	for i := 0; i < len(arr); i += 2 {
		arr[i] = 0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		duplicateZeros(arr)
	}
}

func BenchmarkDuplicateZerosExtraSpace(b *testing.B) {
	arr := make([]int, 10000)
	for i := 0; i < len(arr); i += 2 {
		arr[i] = 0
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		duplicateZerosExtraSpace(arr)
	}
}
