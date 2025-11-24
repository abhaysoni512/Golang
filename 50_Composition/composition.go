package main

import "fmt"

type Engine struct{
	// Engine fields
}

func (e *Engine) Start(){
	fmt.Println("Engine started")
}

type Car struct{
	Engine
	// Car fields
}

func main(){
	my_car := Car{}
	my_car.Start()
}