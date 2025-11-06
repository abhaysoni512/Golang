package main

import (
	"fmt"
	"time"
)

func main()  {
	// simple switcg
	i :=5
	switch i {
	case 1 :
		fmt.Println("1st")
	case 2 :
		fmt.Println("2nd")
	case 3 :
		fmt.Println("3rd")
	case 4 :
		fmt.Println("4th")
	case 5 :
		fmt.Println("5th")
	default:
		fmt.Println("not found")
	}

	// multiple condition switch
	switch time.Now().Weekday(){
	case time.Saturday, time.Sunday:
		fmt.Println("It's weekend")
	default:
		fmt.Println("It's a weekday")
	}
	fmt.Println("----------------------")
	// type switch
	whoAmI := func(i interface{}){
		switch t := i.(type) {        // i.(type) used to return type and storing in i
		case int:
			fmt.Println("its an integer")
		case string:
			fmt.Println("its a string")
		default:
			fmt.Printf("unknown type %T\n", t)
		}
	}

	whoAmI(7)
	whoAmI("hello")
	whoAmI(3.14)
}