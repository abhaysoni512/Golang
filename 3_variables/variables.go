package main

import "fmt"

func main(){

	// 1. we can short it : var name string = "golang"
	var name = "golang"
	fmt.Println(name)

	var age int = 21
	fmt.Println(age)
	
	// 2. one more way to use it short hand way
	name1 := "golang"
	fmt.Println(name1)

	// 3. suppose we want to just declare variable
	var name2 int
	name2 = 11
	fmt.Println(name2)

	var name3 string
	name3 = "hello"
	fmt.Println(name3)


}