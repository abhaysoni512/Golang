package main

import "fmt"

// just as we use template in c++ , here we can also use like that
// we can use any or interface{}
// Requirement : if we want to use only let's say int, string , float data type only then how can we achieve this , we need to use [T int | string | float] instead of it we can use comparable : comparable is an interface that is implemented by all comparable types (booleans, numbers, strings, pointers, channels, arrays of comparable types, structs whose fields are all comparable types). The comparable interface may only be used as a type parameter constraint, not as the type of a variable.

func PrintSlice[T comparable](items[] T){
	for _,item := range items{
		fmt.Println("items : ",item)
	}
}
func main() {
	ex_slice := []int{1,2,3,4,5}
	PrintSlice(ex_slice)

	ex_slice1 := []string{"java","c++","typescript"}
	PrintSlice(ex_slice1)

	// ex_slice2 := []complex128{4+2i,1+3i,2+3i}
	// PrintSlice(ex_slice2)
	// won't be working
}
