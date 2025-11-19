package main

import (
	"fmt"
	"sync"
)

func main() {
	ch := make(chan *int, 4)
	array := []int{1, 2, 3, 4}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer close(ch)
		for _, value := range array {
			v := value
			ch <- &v
		}
	}()
	go func() {
		for value := range ch {
			fmt.Println(*value) // what will be printed here?
		}
		wg.Done()
	}()

	wg.Wait()
}
