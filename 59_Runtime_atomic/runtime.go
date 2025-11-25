package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	fmt.Println("No of Cores : ", runtime.NumCPU())
	fmt.Println("No of Go routines : ", runtime.NumGoroutine())
	var counter int64
	const gs = 100

	var wg sync.WaitGroup
	wg.Add(gs)
	for i := 0; i < gs; i++ {
		go func() {
			defer wg.Done()
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			fmt.Println("Counter\t", atomic.LoadInt64(&counter))
		}()
		fmt.Println("No of Goroutines : ", runtime.NumGoroutine())
	}
	wg.Wait()
	fmt.Println("Hello , playground")
	fmt.Println("Counter : ", counter)

}
