package main

import (
	"context"
	"fmt"
)

func main(){
	ctx := context.Background()
	ctx = context.WithValue(ctx, "UserID", 123)

	done := make(chan struct{})

	go performTask(ctx, done)

	<-done


}

func performTask(ctx context.Context, done chan<- struct{}){
	userId := ctx.Value("UserID")
	fmt.Println(userId)
	done<- struct{}{}
}
