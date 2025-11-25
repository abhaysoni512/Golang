package main

import (
	"fmt"
)


func main() {
	
	c := make(chan int)
	c <- 22

	fmt.Println(<-c)
}
