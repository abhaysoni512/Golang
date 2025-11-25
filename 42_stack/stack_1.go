package main

// s using slices
type Stack interface{
	Push(interface{})
	Pop() (interface{}, bool)
	Peek() (interface{}, bool)
	IsEmpty() bool
	Size() int
}

type SliceStack struct{
	elements[] interface{}
}


func NewSliceStack() *SliceStack{
	return &SliceStack{
		elements: make([]interface{}, 0),
	}
}





//Methods in s : Push, Pop, Peek, Len, IsEmpty

// M0. IsEmpty
func (s * SliceStack) IsEmpty() bool {
	return  len(s.elements) == 0
}

// M1. Push
func (s ) Push(ele interface{}){
	s.elements = append(s.elements, ele)
}

// M2. Pop
func (s *[T]) Pop() (T, bool){
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
func (s *[T]) Peek() (T,bool){
	if s.IsEmpty() {
		var zero T
		return zero, false
	}
	return s.elements[len(s.elements)-1], true
}

func main(){

		

}


