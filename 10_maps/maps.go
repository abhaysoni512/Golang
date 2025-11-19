package main

import (
	"fmt"
)

func main() {
	// Creating map

 m := map[string]string{}

	m["name"] = "golang"
	m["area"] = "backend"

	fmt.Println(m)
	
	fmt.Println(m["area1"]) // it will return empty,0,false, if key not found
	
	
	// m1 := make(map[string]int)
	// fmt.Println(m1["pj"])

	// m1["phone"] = 111111111111
	// m1["score"] = 0222
	// // to delete a

	
	// // delete a key :
	// delete(m1,"score")
	
	// fmt.Println(m1)

	// // to clear map
	// clear(m1)
	// fmt.Println(m1)

	// // declare and initialise value
	// m2 := map[string]int{"price":33,"phone":9800}
	// fmt.Println(m2)
	
	// // to check a key exists on map
	// v,ok := m2["price"] // two things are return (value, true or false)
	// fmt.Println(v)
	
	// if ok {
	// 	fmt.Println("all okay")
	// 	} else{
	// 		fmt.Println("not okay")
			
	// 	}
	// m3 := map[string]int{"price":33,"phone":9800}
	// m4 := map[string]int{"price":33,"phone":9800}

	// fmt.Println(maps.Equal(m3,m4))
		
		



}