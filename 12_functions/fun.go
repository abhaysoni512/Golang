package main

import "fmt"

func add (a int, b int) int{    // func function_name (a int, b int) return_data_type
	return a + b
}

func add1 (a, b int) int{   // if we only use int with last element then by default rest left are int
	return a + b
}

func AddAndSub(a, b int) (int, int){
	return a+b, a-b
}

// In go, functions are first class citizens : can be passed as arguments, can be returned from other functions, can be assigned to variables
func processIt(fn func(a int) int){
	fmt.Println(fn(5))
}

// returning function from another function
func getFunction() func(a int) int {
	return func(a int) int {
		return a * a
	}
}

func main(){
	fmt.Println(add(1,2))
	fmt.Println(add1(1,2))
	fmt.Println(AddAndSub(1,2))

	fmt.Println("-------------------")
	// passing function as argument
	square := func(a int ) int  {
		return a * a
	}
	processIt(square)

	fmt.Println("-------------------")
	// returning function from another function
	sqFunc := getFunction()
	fmt.Println(sqFunc(6))
	
}