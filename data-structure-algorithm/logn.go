package data_structure_algorithm

import (
	"math/rand"
	"sort"
	"time"
)

var testSet = generateIntSliceSorted(1000000)
var randSearchNum = randNumber(testSet)

func generateIntSliceSorted(n int) []int {
	x := make([]int, n)
	for i := 0; i < n; i++ {
		x = append(x, rand.Intn(1000))
	}
	sort.Ints(x)
	return x
}

func randNumber(testSet []int) int {
	rand.Seed(time.Now().Unix())
	randSeed := rand.Int() % len(testSet)
	randSearchNum := testSet[randSeed]
	return randSearchNum
}
func binarySearch(intSlice []int, searchNum int) {
	sort.SearchInts(intSlice, searchNum)
}
func binarySearchTimings(n int) {
	binarySearch(testSet[0:n], randSearchNum)
}
