package main

import (
	"context"
	"fmt"
	"time"
)


func main(){
	ctx, cancel := context.WithCancel(context.Background())

	go performTask(ctx)
	time.Sleep(2 * time.Second)
	cancel()
	time.Sleep(time.Second)
}

func performTask(ctx context.Context ){
	for{
		select{
		case<-ctx.Done():
			fmt.Println("Task canceled ")
			return
		default:
			fmt.Println("Performing task ...")
			time.Sleep(500 * time.Millisecond)
		}
	}
}