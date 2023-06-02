package main

type Ordered interface {
	~float64 | ~int | ~string
}

type StackSlice[T Ordered] struct {
	items []T
}

func getZero[T Ordered]() T {
	var result T
	return result
}

func (stack *StackSlice[T]) Push(item T) {
	if item != getZero[T]() {
		stack.items = append(stack.items, item)
	}
}

func (stack *StackSlice[T]) Pop() T {
	length := len(stack.items)
	if length > 0 {
		returnValue := stack.items[length-1]
		stack.items = stack.items[:(length - 1)]
		return returnValue
	} else {
		return getZero[T]()
	}
}

func (stack StackSlice[T]) IsEmpty() bool {
	return len(stack.items) == 0
}

func (stack StackSlice[T]) Top() T {
	length := len(stack.items)
	if length > 0 {
		return stack.items[length-1]
	} else {
		return getZero[T]()
	}
}
