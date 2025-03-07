package main

import (
	"fmt"
	"log"
	"math/rand"
	"sort"
	"time"
)

func randIntArray(totalGen int) []int {
	data := make([]int, totalGen)
	for i := 0; i < totalGen; i++ {
		t := totalGen
		data[i] = rand.Intn(t)
	}
	sort.Ints(data)
	return data
}
func init() {
	rand.Seed(time.Now().UnixNano())
}

var letterRunes = []rune("abc")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
func randArrStringRunes(n int) []string {
	strArr := make([]string, n)
	for i := 0; i < n; i++ {
		strArr[i] = randStringRunes(5)
	}
	return strArr
}

func main() {
	data := randArrStringRunes(10000)

	result := binarySearchString(data, "aaabc")
	fmt.Println("Result ", result)
}
func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed.Round(time.Microsecond))
}

func binarySearch(data []int, itemTosearch int) bool {
	low, high := 0, len(data)-1
	defer timeTrack(time.Now(), "factorial")

	step := 0
	for low <= high {
		step++
		mid := (low + high) / 2
		guest := data[mid]
		fmt.Println("Guest ", guest, mid, low, high, step)
		if guest == itemTosearch {
			return true
			fmt.Printf("its need %d times\n", step)
		} else if guest > itemTosearch {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return false
}

func binarySearchString(data []string, str string) bool {
	low, high := 0, len(data)-1
	defer timeTrack(time.Now(), "binary search")

	step := 0
	for low <= high {
		step++
		mid := (low + high) / 2
		guest := data[mid]
		fmt.Println("Guest ", guest, mid, low, high, step)
		if guest == str {
			return true
			fmt.Printf("its need %d times\n", step)
		} else if guest > str {
			high = mid - 1
		} else {
			low = mid + 1
		}

	}
	return false
}
