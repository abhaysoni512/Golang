package main

import "fmt"

func main() {
	result1 := add[int](5, 10)
	println("Sum of integers :", result1)
}

// interface function that takes any outputtable data type
func add[T int | string | float32 | float64](a,b T) (T) {
	return a + b
}