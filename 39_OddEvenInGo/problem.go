package main

import "fmt"

// Write program to print Even and Odd numbers using two goroutines
func odd(num int, done chan bool){
	fmt.Println("Odd : ",num)
	done<-true
}

func even(num int, done chan bool){
	fmt.Println("Even : ",num)
	done<-true
}


func main() {
	done := make(chan bool)
	for i := range 100{
		if i%2 != 0{
		  go odd(i,done)
		  <-done
		} else{
		  go even(i,done)
		  <-done
		}
	}
}