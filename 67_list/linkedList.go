package main

import "fmt"

type Node[T comparable] struct {
	data T
	next *Node[T]
}

type Linkedlist[T comparable] struct {
	head   *Node[T]
	length int
}

type List[T any] interface {
	InsertAtHead(val T)
	Traverse()
	Append(val T)
	Delete(val T) bool
}

func (l *Linkedlist[T]) InsertAtHead(val T) {
	newNode := &Node[T]{data: val, next: nil}
	newNode.next = l.head
	l.head = newNode
	l.length++
}

func (l *Linkedlist[T]) Traverse() {
	tmp := l.head
	for tmp != nil {
		fmt.Printf("%v \t", tmp.data)
		tmp = tmp.next
	}
	fmt.Println()
}

func (l *Linkedlist[T]) Append(val T) {
	newNode := &Node[T]{
		data: val,
		next: nil,
	}
	if l.head == nil {
		l.head = newNode
		l.length++
		return
	}

	tmp := l.head
	for tmp.next != nil {
		tmp = tmp.next
	}

	tmp.next = newNode
	l.length++

}

func (l *Linkedlist[T]) Delete(val T) bool {
	if l.head == nil {
		fmt.Println("Empty linkedlist")
		return false
	}

	// C1 : head to be deleted
	if l.head.data == val {
		l.head = l.head.next
		l.length--
		return true
	}

	// c2: need to delete value apart from first
	// 10 20 30 40
	prev := l.head
	curr := l.head.next

	for curr != nil {
		if curr.data == val {
			prev.next = curr.next
			l.length--
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}

func sum[T comparable](tmp *Node[T]) {
	for tmp != nil {

	}
}
func main() {
	// var l Linkedlist[int]

	// l.Append(22)
	// l.Append(33)
	// l.Append(44)

	// sum(l.head)

	var l List[int]

	l = &Linkedlist[int]{}

	l.InsertAtHead(11)
	l.Append(22)
	l.Append(33)
	l.Append(44)
	l.Traverse()

}
