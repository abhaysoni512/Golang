package main

import "fmt"

type OrderStatus int

const (
	Recieved OrderStatus = iota
	Delivered
	Failed
)

type OrderStatus1 string

// we can also give strings here
const (
	Rrecieved  OrderStatus1 = "recived"
	Ddelivered              = "delivered"
	Ffailed                 = "failed"
)

func ChangeOrderStatus(status OrderStatus) {
	fmt.Println("Order status changed to : ", status)
}
func ChangeOrderStatus1(status OrderStatus1) {
	fmt.Println("Order status changed to : ", status)
}

func main() {
	ChangeOrderStatus(Recieved)
	ChangeOrderStatus1(Ddelivered)
	
}
