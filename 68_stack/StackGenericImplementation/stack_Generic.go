package main

import "fmt"

type Stack[T comparable] struct {
	elements []T
}

func newStack[T comparable]() *Stack[T] {
	return &Stack[T]{
		elements: make([]T, 0),
	}
}

func (s *Stack[T]) Push(data T) {
	s.elements = append(s.elements, data)
}

func (s *Stack[T]) IsEmpty() bool {
	return len(s.elements) == 0
}

func (s *Stack[T]) Pop() (T, bool) {
	var zero T
	if s.IsEmpty() {
		return zero, false

	}
	element := s.elements[len(s.elements)-1]
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func (s *Stack[T]) Peek() T {
	var zero T
	if s.IsEmpty() {
		return zero
	}
	return s.elements[len(s.elements)-1]
}

func (s *Stack[T]) Print() {
	if s.IsEmpty() {
		fmt.Println("Stack is empty")
		return
	}
	fmt.Println("Stack elements:\n [ ")
	for i := len(s.elements) - 1; i >= 0; i-- {
		fmt.Printf("  \n%v\n", s.elements[i])
		if i > 0 {
			fmt.Println("_________________")
		}
	}
	fmt.Println(" ]")
}

func main() {
	s := newStack[int]()
	s.Push(11)
	s.Push(22)
	s.Push(33)
	fmt.Println(s.Peek())
	s.Print()

}
