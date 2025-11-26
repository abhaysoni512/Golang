package main

import (
	"fmt"
	"runtime"
)

func main() {
	
	c := make(chan int)

	go foo(c)

	goo(c)

	fmt.Println("No of goroutines : ",runtime.NumGoroutine())

	fmt.Println("About to exit")

}


// send only channel
func foo(c chan<- int){
	for i:=0;i<10;i++{
		fmt.Println("Sender: ",i)
		c<- 42
		
	}
	close(c)
}

// recive only channel
func goo(c <-chan int){
	fmt.Println("No of goroutines : ",runtime.NumGoroutine())
	fmt.Println("Exectued before foo but blocked")
	for v:= range c{
		fmt.Println("Reciever :",v)
	}
}