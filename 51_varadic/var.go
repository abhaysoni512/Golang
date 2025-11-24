package main

import "fmt"

func main(){
	// slice unfluring
	xi := []int{1,2,3,4,5,6,7,8,9}
	x := sum(xi...)

	fmt.Println()

}

func sum(ii ...int) int {
	fmt.Println(ii)
	fmt.Printf("%T\n",ii)
	sum := 0
	for _,v := range ii{

		sum+= v
	}
	return  sum
}