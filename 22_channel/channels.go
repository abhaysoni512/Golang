//Channels are used to communicate between the go routines basically they are used for data tranfer between go routines

package main

import (
	"fmt"
	"math/rand"
	"time"
)
// func processNumber(numchan chan int){
// 	fmt.Println("Go routine processNumber is receving : ",<-numchan)
// }

func processNumber(numchan chan int){
	for{

		fmt.Println("Go routine processNumber is receving : ",<-numchan)
		time.Sleep(time.Second * 1)
	}

}
func main() {
	// messagechannel := make(chan string) // channel creation
	// messagechannel <- "hello"		  // reciving into channel
	// msg := <-messagechannel			  // sending from channel
	// fmt.Println(msg)

	// Above code won't work bcz channel is blocking i.e. channel is blocked means pipe ke jab tk dono end sync m na ho i.e if sender is sending somthing into channel (pipe) reciever must be ready to recieve data

	// numChannel := make(chan int)
	// go processNumber(numChannel) // we should launch go routine before send launches
	// numChannel<-5 // send

	// time.Sleep(time.Microsecond * 2)


	numChannel := make(chan int)
	go processNumber(numChannel) 

	for{
		numChannel<- rand.Intn(100)
	}
	// no need sleep as it's infinte loop
	
}