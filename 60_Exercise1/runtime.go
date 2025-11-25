package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup

func Hi(){
	defer wg.Done()
	fmt.Println("Hi")
}

func Hello(){
	defer wg.Done()
	fmt.Println("Hello")
}

func main() {
	
	wg.Add(2)
	fmt.Println("No of go routines : ", runtime.NumGoroutine())
	go Hi()
	go Hello()

	fmt.Println("No of go routines : ", runtime.NumGoroutine())
	wg.Wait()
}
