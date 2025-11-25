package main

import "fmt"

func main() {
	// create a bidirectional channel and derive unidirectional views
	var ch chan int
	ch = make(chan int, 2)

	send := (chan<- int)(ch) // send-only view
	recv := (<-chan int)(ch) // receive-only view

	// send values
	send <- 22
	send <- 44

	// receive values
	// for v := range recv {
	// 	fmt.Println(v)
	// }
	close(ch) // close the underlying channel

	for i := 0; i < 2; i++ {
		fmt.Println(<-recv)
	}

}
