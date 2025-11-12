package main

import (
	"fmt"
	"sync"
)

func task(id int, w *sync.WaitGroup){
	defer w.Done() // Defer is used to if we want to execute that particular line or block of code once function's body is about to wrap up
	fmt.Println("doing task no : ",id)
}

func main() {

	var wg sync.WaitGroup
	for i:=1;i<=10;i++{
		wg.Add(1)
		go task(i, &wg)
	}

	wg.Wait() // blocks main give time to go routines to finish their task
}