package main

import (
	"fmt"
	"time"
)

const size = 100_000_000

func RunSlice() {
	nameStack := Stack[string]{}
	nameStack.Push("Katy")
	nameStack.Push("Perry")
	nameStack.Push("Harry")
	topStock := nameStack.Top()
	if topStock != getZero[string]() {
		fmt.Printf("\n Top of stack is %s", topStock)
	}
	poppedFromStack := nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	poppedFromStack = nameStack.Pop()
	if poppedFromStack != getZero[string]() {
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}

	intStack := Stack[int]{}
	intStack.Push(5)
	intStack.Push(10)
	intStack.Push(0)
	top := intStack.Top()
	if top != getZero[int]() {
		fmt.Printf("\nValue on top of instStack is %d", top)
	}
	popFromStack := intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped form intStack is %d", popFromStack)
	}
	popFromStack = intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped form intStack is %d", popFromStack)
	}
	popFromStack = intStack.Pop()
	if popFromStack != getZero[int]() {
		fmt.Printf("\nValue popped form intStack is %d", popFromStack)
	}
}

func RunStackNode() {
	nameStack := Stack[string]{}
	nameStack.Push("Lerry")
	nameStack.Push("Page")
	nameStack.Push("Steve")
	if !nameStack.IsEmpty() {
		topOfStack := nameStack.Top()
		fmt.Printf("\nTop of stack is %s", topOfStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	if !nameStack.IsEmpty() {
		poppedFromStack := nameStack.Pop()
		fmt.Printf("\nValue popped from stack is %s", poppedFromStack)
	}
	nameStackInt := Stack[int]{}
	nameStackInt.Push(20)
	nameStackInt.Push(10)
	nameStackInt.Push(20)
	if !nameStackInt.IsEmpty() {
		topOfStack := nameStackInt.Top()
		fmt.Printf("\nTop of stack is %d", topOfStack)
	}
	if !nameStackInt.IsEmpty() {
		poppedFromStack := nameStackInt.Pop()
		fmt.Printf("\nValue popped from stack is %d", poppedFromStack)
	}
	if !nameStackInt.IsEmpty() {
		poppedFromStack := nameStackInt.Pop()
		fmt.Printf("\nValue popped from stack is %d", poppedFromStack)
	}
	if !nameStackInt.IsEmpty() {
		poppedFromStack := nameStackInt.Pop()
		fmt.Printf("\nValue popped from stack is %d", poppedFromStack)
	}
}

func SpeedCompare() {
	nodeStack := StackNode[int]{}
	sliceStack := Stack[int]{}

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
