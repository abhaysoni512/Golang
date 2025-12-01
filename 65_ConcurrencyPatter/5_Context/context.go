package main

import (
	"context"
	"fmt"
	"time"
)

/*
	controlling timeouts, cancelling go rotines and passing some meta data to your Go application
*/

func main(){

}

func exampleTimeOut(){
	ctx := context.Background()

	ctxWithTimeout, cancel := context.WithTimeout(ctx,time.Second *2 )
	defer cancel()

	done := make(chan struct{})
	
	go func(){
		time.Sleep(3 * time.Sleep())
	}()

	select {
	case <-done:
		fmt.Println("Cancelled the API")
	case <- ctxWithTimeout.Done():
		fmt.Println("oH MY TIMEOUT EXPRIED",ctxWithTimeout.Err())
	}
	
}