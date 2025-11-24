package main

import "fmt"


func sliceDelta(ii []int){
	ii[0] = 99
}

func main(){
	xi := []int{1,2,3,4,5}
	fmt.Println(xi)
	
	sliceDelta(xi)
	fmt.Println(xi)

}