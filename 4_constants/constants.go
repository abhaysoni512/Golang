package main

import "fmt"


func main(){
	const name = "golang"

	fmt.Println(name)

	// if we need to give constants to multiple var

	const(
		port = 5000
		host = "localhost"
	)

	fmt.Println(port, host)
}