// here we have for -> only construct for looping in go

package main

import "fmt"

func main(){

	// while loop 
	// i := 1
	// for i<=3 {
	// 	fmt.Println(i)
	// 	i+=1
	// }

	// Infinite loop
	// for{
	// 	fmt.Println(1)
	// }

	// for classic
	// for i :=0; i<3; i++{
	// 	fmt.Println(i)
	// }

	// go 1.22 gives range
	for i:= range 3{
		fmt.Print(i," ")
	}
}