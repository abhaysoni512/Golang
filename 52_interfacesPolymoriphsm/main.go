package main

import "fmt"

type person struct {
	first string
}

type secretAgent struct {
	person
	ltk bool
}

func (p person) speak() {
	fmt.Println("I am ", p.first)
}

func (sa secretAgent) speak() {
	fmt.Println("I am a secret agent", sa.first)
}

type human interface{
	speak()
}

func saySomething( h human){
	h.speak()
}

func main() {
	sa1 := secretAgent{
		person : person{
			first: "jake",
		},
		ltk: true,
	}

	p2 := person{
		first: "jake",
	}

	sa1.speak()
	p2.speak()

	saySomething(sa1)
	saySomething(p2)
}
