package main

import (
	"fmt"
	"time"
	// "math/rand"
	// "time"
	// "sync"
)

// func processNum(nc chan int){
// 	for valInCHan := range nc{
// 		fmt.Println("processing number ",valInCHan)
// 		time.Sleep(time.Second * 1)
// 	}
// }


func emialSender(emailchain chan string,done chan bool){
	defer func(){ done<-true}()
	for email :=range emailchain{
		fmt.Println("sending email to: ",email)
		time.Sleep(time.Second)
	}
}

func main() {
	// msgchan := make(chan string)
	// msgchan<- "ping" // Non - buffered channel is blocking in nature
	// msg := <-msgchan
	// fmt.Println(msg)


	// numChan := make(chan int)
	// go processNum(numChan)

	// for range 5 {
	// 	numChan<- rand.Intn(100)
	// }


	emailchain := make(chan string, 100)
	doneChain := make(chan bool)
	
	go emialSender(emailchain,doneChain)

	for i := range 5{
		emailchain<- fmt.Sprintf("%d@email.com\n",i+1)
	}
	fmt.Println("Done 100 mails in main")
	<-doneChain
}

