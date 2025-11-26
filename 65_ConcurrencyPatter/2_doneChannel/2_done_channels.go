package main

import (
	"fmt"
	"time"
)
// done channel pattern : Here parent goroutine signals to child goroutine to stop its work, it uses a done channel to send the signal.
func doWork(done <-chan bool) {
	for {
		select {
		case <-done:
			fmt.Println("Received done signal, exiting...")
			return
		default:
			fmt.Println("Doing work...")

		}
	}
}
func main() {


	done := make(chan bool)

	go doWork(done)

	for time.Sleep(2 * time.Second); ; {
		done <- true
		break
	}
	// receiving from channel
	

	

}
