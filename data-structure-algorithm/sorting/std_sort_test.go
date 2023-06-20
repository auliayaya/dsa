package sorting

import (
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
)

func TestIsSorted1(t *testing.T) {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	result := isSorted1(data)
	assert.Equal(t, false, result)
	for i := 0; i < size; i++ {
		data[i] = float64(i * 2)
	}
	result = isSorted1(data)
	assert.Equal(t, true, result)
}

func TestIsSorted2(t *testing.T) {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	result := isSorted2(data)
	assert.Equal(t, false, result)
	for i := 0; i < size; i++ {
		data[i] = float64(i * 2)
	}
	result = isSorted2(data)
	assert.Equal(t, true, result)
}

func TestIsSorted3(t *testing.T) {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = 100.0 * rand.Float64()
	}
	result := isSorted3(data)
	assert.Equal(t, false, result)
	for i := 0; i < size; i++ {
		data[i] = float64(i * 2)
	}
	result = isSorted3(data)
	assert.Equal(t, true, result)
}

func BenchmarkIsSorted1(b *testing.B) {
	//data = make([]float64, size)
	//for i := 0; i < size; i++ {
	//	data[i] = float64(i * 2)
	//}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data = make([]float64, size)
		for i := 0; i < size; i++ {
			data[i] = float64(i * 2)
		}
		b.StartTimer()
		isSorted1(data)
	}
}
func BenchmarkIsSorted2(b *testing.B) {
	//data = make([]float64, size)
	//for i := 0; i < size; i++ {
	//	data[i] = float64(i * 2)
	//}
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		data = make([]float64, size)
		for i := 0; i < size; i++ {
			data[i] = float64(i * 2)
		}
		b.StartTimer()
		isSorted2(data)
	}
}

func BenchmarkIsSorted3(b *testing.B) {
	data = make([]float64, size)
	for i := 0; i < size; i++ {
		data[i] = float64(i * 2)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		isSorted3(data)
	}
}
