package main
func Sum(result chan int, num1 int, num2 int){
	result<- num1+num2
}
func main() {

	result := make(chan int)

	go Sum(result, 5 ,4)

	println("sum output : ",<-result)
}