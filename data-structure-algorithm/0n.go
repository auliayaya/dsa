package data_structure_algorithm

func generateIntSlice(n int) []int {
	x := make([]int, 0, n)
	for i := 0; i < n; i++ {
		x = append(x, n)
	}
	return x
}
func simpleLoopSum(count int) int {
	slice := generateIntSlice(count)
	sum := 0
	for i := 0; i < count; i++ {
		sum += slice[i]
	}
	return sum
}
