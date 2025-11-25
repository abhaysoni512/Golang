package main

import (
	"fmt"
	"runtime"
	"sync"
)


func main(){
	fmt.Println("No of Cores : ",runtime.NumCPU())
	fmt.Println("No of Go routines : ",runtime.NumGoroutine())
	counter := 0
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	for i:=0;i<gs;i++{
		go func(){
			defer wg.Done()
			v:= counter
			//time.Sleep(time.Second)
			runtime.Gosched()
			v++
			counter=v
		}()
		fmt.Println("No of Goroutines : ",runtime.NumGoroutine())
	}	
		wg.Wait()
		fmt.Println("Hello , playground")
		fmt.Println("Counter : ", counter)
		
}