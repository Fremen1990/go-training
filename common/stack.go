package common

import "fmt"

type Stack[T any] struct {
	data []T
}

func (s *Stack[T]) Push(element T) {
	s.data = append(s.data, element)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.data) == 0 {
		var empty T
		return empty, false
	}
	lastIndex := len(s.data) - 1
	element := s.data[lastIndex]
	s.data = s.data[:lastIndex]
	return element, true
}

func (s *Stack[T]) Size() int {
	return len(s.data)
}

func testStack() {
	intStack := Stack[int]{}
	intStack.Push(1)
	intStack.Push(2)
	fmt.Println(intStack.Pop())

	floatStack := Stack[float64]{}
	floatStack.Push(1.0)
	floatStack.Push(2.0)
	fmt.Println(floatStack.Pop())

	textStack := Stack[string]{}
	textStack.Push("a")
	textStack.Push("b")
	fmt.Println(textStack.Pop())
}
