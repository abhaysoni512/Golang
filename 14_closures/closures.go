package main

import "fmt"


func counter() func () int{
	count :=0
	return func() int {
		count+=1
		return count
	}
	
}

func main(){

	// Closures : kisi function ke under agr kisi outer scope variable ka use ho rha h na , wo variable humsa available rhga 
	ct := counter()
	fmt.Println(ct())
	fmt.Println(ct())
	
}