package main

// Fan in pattern is used to merge multiple input channels into a single output channel.
// Collecting/combining results from multiple worker goroutines into a single output channel (or reducing them into one value).

import (
	"fmt"
	"sync"
	"time"
)




func main() {
	even := make(chan int)
	odd := make(chan int)
	fanin := make(chan int)

	go send(even, odd)
	go receive(even, odd, fanin)

	for v := range fanin {
		fmt.Println(v)
	}
}

// send channel (producer)

func send(even, odd chan<- int) {
	for i := 0; i < 100; i++ {
		if i%2 == 0 {
			even <- i
		} else {
			odd <- i
		}
	}
	close(even)
	close(odd)
}

// recieve channel

func receive(even, odd <-chan int, fanin chan<- int) {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		for v := range even {
			if v == 50 {
				time.Sleep(time.Second) // artificial delay
			}
			fanin <- v

		}
	}()

	go func() {
		defer wg.Done()
		for v := range odd {
			fanin <- v
		}
	}()

	wg.Wait()
	close(fanin)

}
