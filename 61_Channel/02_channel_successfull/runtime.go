package main

import (
	"fmt"
	"runtime"
)

func main() {
	runtime.GOMAXPROCS(1) // Force single CPU/core
	fmt.Println("No of CPUS : ",runtime.NumCPU())
	c := make(chan int)

	go func() {
		c <- 22
		fmt.Println("sent 22")
	}()
	fmt.Println("No of go routines : ", runtime.NumGoroutine())
	fmt.Println("received:", <-c)
	fmt.Println("No of go routines : ", runtime.NumGoroutine())
}	
