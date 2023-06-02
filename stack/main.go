package main

import (
	"fmt"
	"time"
)

const size = 100_000_000

func RunSlice() {
	nameStackSlice := StackSlice[string]{}
	nameStackSlice.Push("Katy")
	nameStackSlice.Push("Perry")
	nameStackSlice.Push("Harry")
	topStock := nameStackSlice.Top()
	if topStock != getZero[string]() {
		fmt.Printf("\n Top of stack is %s", topStock)
	}
	poppedFromStackSlice := nameStackSlice.Pop()
	if poppedFromStackSlice != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStackSlice)
	}
	poppedFromStackSlice = nameStackSlice.Pop()
	if poppedFromStackSlice != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStackSlice)
	}
	poppedFromStackSlice = nameStackSlice.Pop()
	if poppedFromStackSlice != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStackSlice)
	}

	intStackSlice := StackSlice[int]{}
	intStackSlice.Push(5)
	intStackSlice.Push(10)
	intStackSlice.Push(0)
	top := intStackSlice.Top()
	if top != getZero[int]() {
		fmt.Printf("\nValue on top of instStackSlice is %d", top)
	}
	popFromStackSlice := intStackSlice.Pop()
	if popFromStackSlice != getZero[int]() {
		fmt.Printf("\nValue popped form intStackSlice is %d", popFromStackSlice)
	}
	popFromStackSlice = intStackSlice.Pop()
	if popFromStackSlice != getZero[int]() {
		fmt.Printf("\nValue popped form intStackSlice is %d", popFromStackSlice)
	}
	popFromStackSlice = intStackSlice.Pop()
	if popFromStackSlice != getZero[int]() {
		fmt.Printf("\nValue popped form intStack is %d", popFromStackSlice)
	}
}

func RunStackNode() {
	nameStackNode := StackNode[string]{}
	nameStackNode.Push("Lerry")
	nameStackNode.Push("Page")
	nameStackNode.Push("Steve")
	if !nameStackNode.IsEmpty() {
		topOfStackNode := nameStackNode.Top()
		fmt.Printf("\nTop of stack is %s", topOfStackNode)
	}
	if !nameStackNode.IsEmpty() {
		poppedFromStackNode := nameStackNode.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStackNode)
	}
	if !nameStackNode.IsEmpty() {
		poppedFromStackNode := nameStackNode.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStackNode)
	}
	if !nameStackNode.IsEmpty() {
		poppedFromStackNode := nameStackNode.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStackNode)
	}
	nameStackNodeInt := StackNode[int]{}
	nameStackNodeInt.Push(20)
	nameStackNodeInt.Push(10)
	nameStackNodeInt.Push(20)
	if !nameStackNodeInt.IsEmpty() {
		topOfStackNode := nameStackNodeInt.Top()
		fmt.Printf("\nTop of stack is %d", topOfStackNode)
	}
	if !nameStackNodeInt.IsEmpty() {
		poppedFromStackNode := nameStackNodeInt.Pop()
		fmt.Printf("\nValue popped from stack is %d", poppedFromStackNode)
	}
	if !nameStackNodeInt.IsEmpty() {
		poppedFromStackNode := nameStackNodeInt.Pop()
		fmt.Printf("\nValue popped from stack is %d", poppedFromStackNode)
	}
	if !nameStackNodeInt.IsEmpty() {
		poppedFromStackNode := nameStackNodeInt.Pop()
		fmt.Printf("\nValue popped from stack is %d", poppedFromStackNode)
	}
}

func SpeedCompare() {
	nodeStack := StackNode[int]{}
	sliceStack := StackSlice[int]{}

	start := time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Push(i)
	}
	elapsed := time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on nodeStack: ", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		nodeStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on nodeStack: ", elapsed)

	// Slice Stack
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Push(i)
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Push() operations on sliceStack: ", elapsed)
	start = time.Now()
	for i := 0; i < size; i++ {
		sliceStack.Pop()
	}
	elapsed = time.Since(start)
	fmt.Println("\nTime for 10 million Pop() operations on sliceStack: ", elapsed)

}
func main() {
	//RunSlice()
	//RunStackNode()
	SpeedCompare()
}
