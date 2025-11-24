package main

import (
	"log"
	"os"
)

/*
	One of the most ubiquitous interfaces is Stringer defined by the fmt package.

	type Stringer interface {
		String() string
	}
	A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.


*/



func main() {
	f, err := os.Create("output.txt")
	if err != nil{
		log.Fatalf("error %s",err)
	}

	defer f.Close()

	s := []byte("Hello Golang")

	_, err = f.Write(s)
	if err != nil{
		log.Fatalf("error %s",err)
	}


}
