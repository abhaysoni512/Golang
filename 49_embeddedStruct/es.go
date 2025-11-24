package main

import "fmt"

type person struct {
	firstName string
	lastName  string
	age       int
}

type secretAgent struct {
	person
	ltk bool
}

func main() {
	sa1 := secretAgent{
		person: person{
			firstName: "James",
			lastName:  "Bond",
			age:       29,
		},
		ltk: true,
	}

	fmt.Println("First name : ", sa1.person, "ltk : ", sa1.ltk)
	fmt.Printf(" p1 is of type %T\n",sa1)

	fmt.Println("-------------------------")

	p1 := struct {
		firstName string
		lastName  string
		age       int
	}{
		firstName: "jAME",
		lastName:  "Bond",
		age:       29,
	}

	fmt.Printf(" p1 is of type %T\n",p1)

}
