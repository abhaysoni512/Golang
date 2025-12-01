package main

import "fmt"
type Stack struct{
	elements[] any
}

func newStack() *Stack{
	return &Stack{
		elements: make([]any, 0),
	}
}

func (s *Stack)Push(data any){
	s.elements = append(s.elements, data)
}

func (s * Stack)IsEmpty() bool{
	return len(s.elements) == 0
}

func (s *Stack)Pop() (any, bool){
	var zero any
	if s.IsEmpty(){
		return zero, false
	}
	l := len(s.elements)-1
	val := s.elements[l]
	s.elements = s.elements[:l]
	return val, true

}

func (s *Stack) Print(){
	if s.IsEmpty(){
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
func main(){
	s := newStack()
	s.Push(11)
	s.Push("Abhay")
	s.Push(15)
	s.Push(20)
	s.Push(22)
	fmt.Println(s.Pop())
	s.Print()
}