package main

import (
	"fmt"
	"log"
	"strconv"
)

/*
	One of the most ubiquitous interfaces is Stringer defined by the fmt package.

	type Stringer interface {
		String() string
	}
	A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.


*/

type book struct{
	title string
}

func (b book) String() string{
	return fmt.Sprint("the book title is : ",b.title)
}

type count int 
func (c count) String() string{
	return fmt.Sprint("the number is ", strconv.Itoa(int(c)))
}

func logInfo(s fmt.Stringer){
	log.Println("LOG FROM 138", s.String())
}



func main() {
	my_book := book{
		title: "The man who knows infinity",
	}

	logInfo(my_book)

	var num count = 11
	
	logInfo(num)


}
