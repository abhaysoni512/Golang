package main

import "fmt"

func main() {
	userNames := make([]string,2)
	userNames[0] = "1st User"
	userNames[1] = "2nd User"
	fmt.Println(userNames,len(userNames))

	userNames = append(userNames, "Abhay")
	userNames = append(userNames, "John")
	userNames = append(userNames, "Doe")
	fmt.Println(userNames,len(userNames))


	courseRatings := make(map[string]float64, 5)

	courseRatings["ReactJS"] = 4.5
	courseRatings["Golang"] = 4.8
	courseRatings["Docker"] = 4.3
	
	fmt.Println(courseRatings,len(courseRatings))
}