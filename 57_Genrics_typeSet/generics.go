package main

import "fmt"

// typeSet
// ~: include all values of type int, float64 , also underlying type alias
type myDatatypes interface{
	~int | ~float64
}

func addT[T myDatatypes](a, b T) T {
	return a + b
}
 
type myAlias int

func main(){	
	var n myAlias = 42
	fmt.Println(addT(1,2))
	fmt.Println(addT(n,2))
}