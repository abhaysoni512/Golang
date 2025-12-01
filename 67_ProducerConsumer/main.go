package main

import (
	"fmt"
	"time"
)

func producer(ch chan<- int) {
	defer func() {
		close(ch)

	}()
	for i := 0; i < 10; i++ {
		fmt.Printf("Producer : Sending %v\n", i)
		ch <- i
		time.Sleep(100 * time.Microsecond)
	}

	fmt.Println("✅ Producer finished and closed channel.")
}

func consumer(ch <-chan int, done chan<- struct{}) {
	defer close(done)
	for item := range ch {
		fmt.Printf("⬅️ Consumer: Received %d\n", item)
		time.Sleep(500 * time.Millisecond) // Simulate slow work
	}

	fmt.Println("✅ Consumer finished processing all items.")
}

func main() {

	done := make(chan struct{})

	ch := make(chan int, 5)

	go producer(ch)
	go consumer(ch, done)
	fmt.Println("Main is waiting for go routines to finish their task...")
	<-done
	fmt.Println("main done...")
}
