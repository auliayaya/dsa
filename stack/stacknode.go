package main

type Node[T any] struct {
	value T
	next  *Node[T]
}

type StackNode[T any] struct {
	first *Node[T]
}

func (stack *StackNode[T]) Push(item T) {
	newNode := Node[T]{item, nil}
	newNode.next = stack.first
	stack.first = &newNode
}

func (stack *StackNode[T]) Top() T {
	return stack.first.value
}

func (stack *StackNode[T]) Pop() T {
	result := stack.first.value
	stack.first = stack.first.next
	return result
}

func (stack StackNode[T]) IsEmpty() bool {
	return stack.first == nil
}
