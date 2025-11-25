package main

import "fmt"

type dog struct{
	first string
}

func (d dog) walk(){
	fmt.Println("My name is : ",d.first," and i am walking.")	
}
func (d *dog) run(){
	d.first = "Rover"
	fmt.Println("My name is : ",d.first," and i am running.")	
}
type younin interface{
	walk()
	run()
}

func youngRun(y younin){
	y.run()
}
func main(){
	d1 := dog{"Henry"}
	d1.walk()
	d1.run()
	// youngRun(d1)

	d2 := &dog{"padget"}
	d2.walk()
	d2.run()
	youngRun(d2)

	

}