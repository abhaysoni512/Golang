package main

// by val
func changeValue(num int) {
	num =10
	println("num inside changeValue : ",num)
} 
// by reference

func changeValueByReference(num *int){
	*num = 10
	println("num inside changeValueByReference : ", *num)

}
func main()  {
	
	//num := 1
	//changeValue(num)
	//println("num inside main : ",num)

	num :=1
	changeValueByReference(&num)
	println("num inside main : ",num)
}
