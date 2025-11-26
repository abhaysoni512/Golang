package main

import (
	"fmt"
)

func sliceToChannel(nums []int) <-chan int {
	out := make(chan int)

	go func() {

		for _, val := range nums {
			out <- val
		}
		close(out)
	}()
	
	return out
}

// stage 2: square the numbers from input channel and send to output channel(read only)

func sq(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		for n := range in {
			out <- n * n
		}
		close(out)
	}()
	return out
}
func main() {
	// input
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// stage 1 convert data of nums slice to channel
	dataChannel := sliceToChannel(nums)

	// stage 2
	finalChannel := sq(dataChannel)

	// stage 3 read the content from final channel
	for val := range finalChannel {
		fmt.Println(val)
	}

}
