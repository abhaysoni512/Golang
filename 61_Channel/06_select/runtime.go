package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	even := make(chan int)
	odd := make(chan int)
	done := make(chan int)

	go send(even, odd, done)
	go func() {
		receive(even, odd, done)
		wg.Done()
	}()

	wg.Wait()

}

func receive(e, o, q <-chan int) {
	for {
		select {
		case v := <-e:
			fmt.Println("From even channel :", v)

		case v := <-o:
			fmt.Println("From odd channel :", v)
		case v ,ok := <-q :
			if !ok{
			fmt.Println("From quit channel :", v)
			return
			} else{
				fmt.Println("Channel isn't close from sender side")
				panic(v)
			}
		}

	}
}

func send(e, o, q chan<- int) {

	defer func() {
		close(q)
	}()

	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			e <- i
		} else {
			o <- i
		}
	}
	

}
