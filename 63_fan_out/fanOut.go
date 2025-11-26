package main

import (
	"fmt"
	"sync"
	"time"
)

// The Fan-Out design pattern is a strategy used in concurrent programming to distribute work efficiently. It involves a single producer (or a small group of producers) that sends data/tasks to multiple workers (Goroutines) that process the data simultaneously.
// Core Concept:

// Input Channel (Producer): A single channel acts as the source, where the producer Goroutine writes the data or tasks to be processed.

// Multiple Goroutines (Workers): Several worker Goroutines read from this same input channel.

// Concurrent Processing: Because multiple workers are consuming from the same source, the work is processed in parallel, dramatically increasing throughput.

// The total number of tasks to distribute
const numJobs = 10

// The number of workers to fan out to
const numWorkers = 2

func main() {
	input := make(chan int) // input channel
	var wg sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg.Add(1)
		go worker(w, input, &wg)
	}

	go sendJobs(input)

	wg.Wait()
	fmt.Println("All workers done")
}

func sendJobs(input chan<- int) {

	defer close(input) // Always close the channel when production is done

	fmt.Println("Producer: Sending 10 jobs...")
	for i := 1; i <= numJobs; i++ {
		input <- i
	}

}

func worker(id int, input <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for job := range input {
		fmt.Printf("Worker %d: Started job %d\n", id, job)
		time.Sleep(time.Second) // Simulate time-consuming work
		fmt.Printf("Worker %d: Finished job %d\n", id, job)
	}
}
