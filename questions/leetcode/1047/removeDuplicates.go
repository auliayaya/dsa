package leetcode

func removeDuplicateWithStack(S string) string {
	var stack []byte
	stack = append(stack, S[0])
	for i := 0; i < len(S); i++ {
		if len(stack) == 0 || S[i] != stack[len(stack)-1] {
			stack = append(stack, S[i])
		} else {
			stack = stack[:len(stack)-1]
		}
	}
	return string(stack)
}

func removeDuplicateTwoPointer(S string) string {
	bytes := []byte(S)
	var end int
	for _, c := range bytes {
		bytes[end] = c
		if end > 0 && bytes[end-1] == c {
			end -= 2
		}
		end++
	}
	return string(bytes[:end])
}
