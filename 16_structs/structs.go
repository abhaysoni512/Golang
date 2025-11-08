package main

import (
	"fmt"
	"time"
)
type custmer struct{
	name string
	phone_no string
}

type paymentOrder struct{
	name string
	amount float32
	status string
	time time.Time
	custmer
}
// Constructor Work Arroud 
func newOrder(name string,amount float32, status string) (*paymentOrder){
	order := paymentOrder{
		name: name,
		amount: 122.22,
		status: "successful",
	}
	return &order
}

// If we want to change behaviour , we need to use method.
// in order to connect method with structure
// Use : reciever type()
func (p *paymentOrder)changeStatus(status string){
	p.status = status

}

func (p paymentOrder)getstatus() string{
	return p.status
}
func main(){
	// order := paymentOrder{
	// 	name: "phone",
	// 	amount: 122.22,
	// 	status: "successful",
	// }
	// order.time = time.Now()

	// fmt.Println(order)
	// fmt.Println(order.name)



	// Behaviour Utilization
	// order := paymentOrder{
	// 	name: "phone",
	// 	amount: 122.22,
	// 	status: "successful",
	// }
	// order.time = time.Now()
	
	// order.changeStatus("delivered")
	// fmt.Println(order)
	//fmt.Println(order.getstatus())


	// Constructor usage : 
	// myOrder := newOrder("car",100, "delivered")
	// fmt.Println(myOrder)

	// fmt.Println(myOrder.name)

	// if we just need use structure once a time we can define struct in one place
	// language :=struct {
	// 	name string
	// 	isGood bool
	// }{"golang",true}
	// fmt.Println(language)
	newCustmer := custmer{
		name: "John",
		phone_no: "1234567890",
	}
	myOrder := paymentOrder{
		name: "pager",
		amount: 20,
		status: "Delivered",
		time: time.Now(),
		custmer: newCustmer,
	}
	fmt.Println(myOrder)
}
