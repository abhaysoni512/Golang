package main

import "fmt"

type payment struct {
	Gateway paymentGateway
}

// We will use interface it's act as contract between that jo bhi struct ki instance method use kr rhi h use ek method implement jror krni pdegi
type paymentGateway interface {
	pay(amount float32)
}

// as we modify Gateway, so it's voilating open close priciple i.e. your classes and methods should be open for new extension but closed for modification for existing one
func (p payment) makePayment(amount float32) {
	// Instead of doing this we will utilize interface
	// razorPayGateWay := razorPay{}
	// razorPayGateWay.pay(amount)
	// stripPayGateway := stripPay{}
	// stripPayGateway.pay(amount)
	p.Gateway.pay(amount)

}

// Now which Gateway we need to use :
type razorPay struct{}

func (r razorPay) pay(amount float32) {
	fmt.Println("Making a payment through razorPay with amount : ", amount)
}

type stripPay struct{}

func (s stripPay) pay(amount float32) {
	fmt.Println("Making a payment through stripe with amount : ", amount)
}

func main() {

	//newPayment := payment{Gateway: razorPay{}}
	newPayment := payment{Gateway: stripPay{}}
	newPayment.makePayment(100)
}
