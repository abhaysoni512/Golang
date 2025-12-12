package main

import (
	"fmt"
)


func main(){
	nums := []int{1,2,3,4,5,1,2,2}

	// Find duplicate int count and top K duplicate ints
	FrequencyChecker := make(map[int]int)

	for _,val := range nums{
		FrequencyChecker[val]++
	}

	duplicateIntCounter := 0
	kIntgers := 0
	for _, freq := range FrequencyChecker{
		if freq > 1 {
			duplicateIntCounter++
		}
	}
	fmt.Println(duplicateIntCounter)


}