package main 

import "fmt"

type Queue struct{
	Elements []int
	Size int
}

func (q *Queue) Enqueue(elem int){
	if len(q.Elements) == q.Size{
		fmt.Println("Queue is full")
		return
	}
	q.Elements = append(q.Elements, elem)
}

func (q *Queue) Dequeue(elem int){
	if len(q.Elements) == 0 {
		fmt.Println("Queue is empty")
		return
	}

	q.Elements = q.Elements[1:]
}



func main(){

	queue := Queue{Size: 10}

	
}
