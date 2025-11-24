package main

import "fmt"

func main(){
	x := doMath(10,2,add)
	fmt.Println(x)

	y := doMath(10,2,subtract)
	fmt.Println(y)
}

func doMath(a,b int, f func (int,int) int) int{
	return f(a,b)
}

func add(a int, b int)int {
	return  a+b
}

func subtract(a int, b int)int {
	return  a-b
}