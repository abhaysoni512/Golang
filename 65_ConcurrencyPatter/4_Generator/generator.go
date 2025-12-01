package main

import (
	"fmt"
	"math/rand"
)

// it will generate infinte stream of data for channel based on some function passed to it
func repeatFunc[T any, K any](done <-chan K, fn func() T) <-chan T {

	stream := make(chan T)
	go func() {
		for {
			select {
			case <-done:
				return
			case stream <- fn():
			}
		}
	}()
	return stream
}

func take[T any, K any](done <-chan K, stream <-chan T, n int) <-chan T {
	taken := make(chan T)
	go func() {
		defer close(taken)
		for i := 0; i < n; i++ {
			select {
			case <-done:
				return
			case taken <- <-stream:
			}
		}
	}()
	return taken
}

func main() {
	done := make(chan int)
	defer close(done)

	randNumFetcher := func() int { return rand.Intn(100) }

	// for randNo := range repeatFunc(done, randNumFetcher) {
	// 	fmt.Println(randNo)
	// }
	randIntStream := repeatFunc(done, randNumFetcher)
	for randno := range take(done, randIntStream, 10) {
		fmt.Println(randno)
	}
}
