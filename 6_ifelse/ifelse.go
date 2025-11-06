package main

import "fmt"

func main(){
	//age := 18
	// if age>=18 {
	// 	fmt.Println("Person is an adult")
	// } else{	// here and if block line } else must be on single line
	// 	fmt.Println("Person is an adult")
	// }

	var role = "admin"
	var hasPermission = false

	if role == "admin" && hasPermission {
		fmt.Println("yes")
	}

	// go doesn't have terniary operator 1.25
}
