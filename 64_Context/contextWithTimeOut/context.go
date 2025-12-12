// Context provides a mechanism to control the lifecycle, cancellation, and propagation of requests across multiple goroutines. With context, you can create a hierarchy of goroutines and pass important information down the chain. (CPrL )

package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	done := make(chan string)

	go performTask(ctx,done)

	select {
	case msg := <-done:
		fmt.Println(msg)
	case <-ctx.Done():
		fmt.Println("Task was timed out (or canceled)")
	}

}

func performTask(ctx context.Context, done chan<- string) {
	select {
	case <-time.After(5 * time.Second):
		fmt.Println("Task completed successfully")
	case <-ctx.Done():
		done <- "Task was time out : " + ctx.Err().Error()
	}
}
