package main

import "fmt"

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{elements: make([]T, 0)}
}

func (s *Stack[T]) Push(value T) {
	s.elements = append(s.elements, value)
}
func (s *Stack[T]) Pop() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	retVal := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return retVal, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func main() {
	stack := NewStack[int]()
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())
	for i := range 10 {
		stack.Push(i)
	}
	fmt.Println(stack.Peek())
	fmt.Println(stack.IsEmpty())
	fmt.Println(stack.Pop())
	fmt.Println(stack.Peek())

	type CustomType struct {
		i int
		b bool
		s string
	}

	stack2 := NewStack[CustomType]()
	stack2.Push(CustomType{i: 10, b: true, s: "hello"})
	fmt.Println(stack2.Peek())
	fmt.Println(stack2.IsEmpty())
	fmt.Println(stack2.Pop())
}
