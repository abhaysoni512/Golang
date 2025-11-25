package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func Hi() {
	defer wg.Done()
	fmt.Println("Hi")
}

func Hello() {
	defer wg.Done()
	fmt.Println("Hello")
}

func main() {

}
