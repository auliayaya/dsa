package sorting

import (
	"runtime"
	"sort"
)

const size = 100_000_000

var data []float64

func isSorted1(data []float64) bool {
	var data1 []float64
	data1 = make([]float64, len(data))
	copy(data1, data)
	sort.Float64s(data1)
	for i := 0; i < len(data1); i++ {
		if data[i] != data1[i] {
			return false
		}
	}
	return true
}

func isSorted2(data []float64) bool {
	for i := 1; i < len(data); i++ {
		if data[i] < data[i-1] {
			return false
		}
	}
	return true
}

func isSegmentSorted(data []float64, a, b int, ch chan<- bool) {
	for i := a + 1; i < b; i++ {
		if data[i] < data[i-1] {
			ch <- false
		}
	}
	ch <- true
}
func isSorted3(data []float64) bool {
	ch := make(chan bool)
	numSegments := runtime.NumCPU()
	segmentSize := int(float64(len(data)) / float64(numSegments))
	for index := 0; index < numSegments; index++ {
		go isSegmentSorted(data, index*segmentSize, index*segmentSize+segmentSize, ch)
	}
	num := 0
	for {
		select {
		case value := <-ch:
			if value == false {
				return false
			}
			num += 1
			if num == numSegments {
				return true
			}
		}
		return true
	}
}
