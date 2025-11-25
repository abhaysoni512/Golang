package main

import (
	"fmt"
	"math"
)

type Circle struct{
	radius float64
}

func (c *Circle) Area() float64{
	return  math.Pi * c.radius * c.radius
}


type Shape interface{
	Area() float64
}

func ShapeInfo(s Shape){
	fmt.Println("Area : ", s.Area())
}

func main(){
	c := Circle{11.2}
	c.Area()

	// Method set of c value of TYPE C is phi(empty ) as reciver required of TYPE C in method isn't here
	
	
	
}
