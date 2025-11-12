/*
	GoRoutines : These are light weight threads which are used to achieve multithreading, concurrency
				 With Go Routines we can have background workers, 
*/

package main

import (
	"fmt"
	"time"
)
func task (id int){
	fmt.Println("Doing task with : ",id)
}
func main() {
	for i := range 10 {
		go task(i)
	}

	time.Sleep(time.Millisecond * 2)
}