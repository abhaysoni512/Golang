package main

import "fmt"

// go routine synchroniser
func task(done chan bool) {
	defer func() { done <- true }() // defer func(){t<-true}  function defenition  () will call it
	fmt.Println("Processing")
}

func main() {
	// It's unbuffered channel means sending and recieving is blocking, means ek hi time m either send ya either recive hoga
	done := make(chan bool)

	go task(done)
	<-done
}
