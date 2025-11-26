package main

import "fmt"

func main() {
	
	c := make(chan int)

	go foo(c)

	goo(c)

	fmt.Println("About to exit")

}


// send only channel
func foo(c chan<- int){
	c<- 42
}

// recive only channel
func goo(c <-chan int){
	fmt.Println("Exectued before foo but blocked")
	fmt.Println(<-c)
}