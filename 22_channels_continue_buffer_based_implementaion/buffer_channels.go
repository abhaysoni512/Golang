package main

import (
	"fmt"
	"time"
)
// type safety in channel means if we want to enusre that particular go routine we only want to send in channel or recieve in channel, we will use it using <- operation : (<-chan) ensures we can only recieve data from channel 
// chan<- ensures that ki channel me data sirf send kr skte h

func emailSender(emailchain <-chan string, done chan<- bool) {
	defer func() { done <- true }()
	for email := range emailchain {
		fmt.Println("Email sending from worker function: ", email) // np need to add <- on range case
		time.Sleep(time.Second)
	}
}

func main() {
	//buffer based channel means we can send data of 100 byte after that we will get blocker one
	emailchain := make(chan string, 10)
	done := make(chan bool)
	go emailSender(emailchain, done)
	for i := range 10 {
		emailchain <- fmt.Sprintf("%d@email.com", i+1)
	}
	close(emailchain)
	fmt.Println("Email processing done from main ")

	<-done // will block main routine until go routines get complmented
}
