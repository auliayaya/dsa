package main

import "strconv"

type Element struct {
	value int
}

func (s *Element) String() string {
	return strconv.Itoa(s.value)
}

type Stack struct {
	elements []Element
}

//func (s *Stack) New() {
//	s.elements = make(Elem , 0)
//}
