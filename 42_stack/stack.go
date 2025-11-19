package main

import "fmt"

// Stacks using slices
type Stack[T any] struct{
	elements[] T
}

//Methods in stacks : Push, Pop, Peek, Len, IsEmpty

// M0. IsEmpty
func (s *Stack[T]) IsEmpty() bool {
	return  len(s.elements) == 0
}

// M1. Push
func (s *Stack[T]) Push(ele T){
	s.elements = append(s.elements, ele)
}

// M2. Pop
func (s *Stack[T]) Pop() (T, bool){
	if s.IsEmpty() == true{
		var zero T
		return zero,false
	}
	new_index := len(s.elements) - 1
	top_elem := s.elements[new_index]
	s.elements = s.elements[:new_index]
	return top_elem, true
	

}

// M3. Peek
func (s *Stack[T]) Peek() (T,bool){
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

func main(){
	s := Stack[int]{}

	s.Push(1000)
	s.Push(11)
	s.Push(111)

	if val,ok := s.Peek() ;ok{
		fmt.Println("top element : ",val)
	}

	for !s.IsEmpty(){
		if val,ok := s.Pop() ; ok{
			fmt.Println(val)
		}
	}

	
	type Student struct{
		name string
		age int
		
	}

	stStack := Stack[Student]{}

	stStack.Push(Student{"Abhay",11})
	stStack.Push(Student{"Alex",22})

	if val,ok := stStack.Peek() ;ok{
		fmt.Println("top element : ",val.age,val.name)
	}

		

}


